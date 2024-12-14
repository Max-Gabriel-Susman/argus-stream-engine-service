package main

import (
	"github.com/Max-Gabriel-Susman/argus-stream-engine-service/internal/rtmp"
)

func main() {
	rtmp.LogInfo("Initializing Argus Stream Engine...")

	rtmpServer := rtmp.NewServer()
	if rtmpServer != nil {
		rtmpServer.Start()
	}

	rtmp.LogInfo("Argus Stream Engine Online.")
}
