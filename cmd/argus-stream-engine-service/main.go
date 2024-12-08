package main

import (
	"github.com/Max-Gabriel-Susman/argus-stream-engine-service/internal/hls"
	"github.com/Max-Gabriel-Susman/argus-stream-engine-service/internal/logging"
	"github.com/Max-Gabriel-Susman/argus-stream-engine-service/internal/redis"
	"github.com/Max-Gabriel-Susman/argus-stream-engine-service/internal/rtmp"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	logging.LogInfo("Initializing Argus Stream Engine...")

	rtmpServer := rtmp.CreateRTMPServer()

	go redis.SetupRedisCommandReceiver(rtmpServer)

	if rtmpServer != nil {
		rtmpServer.Start()
	}

	hlsServer := hls.NewServer("./videos")
	hlsServer.ServeContent()

	logging.LogInfo("Argus Stream Engine Online.")
}
