package main

import (
	"github.com/Max-Gabriel-Susman/argus-stream-engine-service/internal/logging"
	"github.com/Max-Gabriel-Susman/argus-stream-engine-service/internal/rtmp"
)

func main() {
	logging.LogInfo("Initializing Argus Stream Engine...")

	rtmpServer := rtmp.NewServer()
	if rtmpServer != nil {
		rtmpServer.Start()
	}

	logging.LogInfo("Argus Stream Engine Online.")
}
