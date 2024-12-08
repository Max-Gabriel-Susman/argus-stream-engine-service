package redis

import (
	"context"
	"crypto/tls"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/Max-Gabriel-Susman/argus-stream-engine-service/internal/logging"
	"github.com/Max-Gabriel-Susman/argus-stream-engine-service/internal/rtmp"
	"github.com/redis/go-redis/v9"
)

func SetupRedisCommandReceiver(server *rtmp.RTMPServer) {
	useRedis := os.Getenv("REDIS_USE")

	if useRedis != "YES" {
		return // Not using redis
	}

	defer func() {
		if err := recover(); err != nil {
			switch x := err.(type) {
			case string:
				logging.LogError(errors.New(x))
			case error:
				logging.LogError(x)
			default:
				logging.LogError(errors.New("could not connect to redis"))
			}
		}
		logging.LogWarning("Connection to Redis lost!")
	}()

	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "localhost"
	}

	redisPort := os.Getenv("REDIS_PORT")
	if redisPort == "" {
		redisPort = "6379"
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisChannel := os.Getenv("REDIS_CHANNEL")

	if redisChannel == "" {
		redisChannel = "rtmp_commands"
	}

	redisTLS := os.Getenv("REDIS_TLS")

	ctx := context.Background()

	var redisClient *redis.Client

	if redisTLS == "YES" {
		redisClient = redis.NewClient(&redis.Options{
			Addr:      redisHost + ":" + redisPort,
			Password:  redisPassword,
			TLSConfig: &tls.Config{},
		})
	} else {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     redisHost + ":" + redisPort,
			Password: redisPassword,
		})
	}

	subscriber := redisClient.Subscribe(ctx, redisChannel)

	logging.LogInfo("[REDIS] Listening for commands on channel '" + redisChannel + "'")

	for {
		msg, err := subscriber.ReceiveMessage(ctx)

		if err != nil {
			logging.LogWarning("Could not connect to Redis: " + err.Error())
			time.Sleep(10 * time.Second)
		} else {
			// Parse message
			parseRedisCommand(server, msg.Payload)
		}
	}
}

func parseRedisCommand(server *rtmp.RTMPServer, cmd string) {
	defer func() {
		if err := recover(); err != nil {
			switch x := err.(type) {
			case string:
				logging.LogError(errors.New(x))
			case error:
				logging.LogError(x)
			default:
				logging.LogError(errors.New("parsing error"))
			}
		}
		logging.LogWarning("Could not parse message: " + cmd)
	}()

	parts := strings.Split(cmd, ">")
	if len(parts) != 2 {
		logging.LogWarning("Invalid message from Redis: " + cmd)
		return // Invalid message
	}

	cmdName := parts[0]
	cmdArgs := strings.Split(parts[1], "|")

	switch cmdName {
	case "kill-session":
		if len(cmdArgs) < 1 {
			logging.LogWarning("Invalid message from Redis: " + cmd)
			return
		}
		channel := cmdArgs[0]
		publisher := server.GetPublisher(channel)

		if publisher != nil {
			publisher.Kill()
		}
	case "close-stream":
		if len(cmdArgs) < 2 {
			logging.LogWarning("Invalid message from Redis: " + cmd)
			return
		}

		channel := cmdArgs[0]
		streamId := cmdArgs[1]
		publisher := server.GetPublisher(channel)

		if publisher != nil && publisher.StreamID == streamId {
			publisher.Kill()
		}
	default:
		logging.LogWarning("Unknown Redis command: " + cmd)
	}
}
