package main

import (
	"github.com/Max-Gabriel-Susman/argus-stream-engine-service/internal/logging"
	"github.com/Max-Gabriel-Susman/argus-stream-engine-service/internal/redis"
	"github.com/Max-Gabriel-Susman/argus-stream-engine-service/internal/rtmp"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	logging.LogInfo("RTMP Go Server (Version 1.0.0)")

	server := rtmp.CreateRTMPServer()

	go redis.SetupRedisCommandReceiver(server)

	if server != nil {
		server.Start()
	}
}
