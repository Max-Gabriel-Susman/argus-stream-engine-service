package main

import (
	"fmt"

	"github.com/Max-Gabriel-Susman/argus-stream-engine-service/internal/detection"
	"github.com/Max-Gabriel-Susman/argus-stream-engine-service/internal/server"
	"github.com/Max-Gabriel-Susman/argus-stream-engine-service/internal/stream"
)

func main() {
	run()
}

func run() error {
	server, err := server.NewStreamServer(stream.NewPipeline(), detection.NewDetector())
	if err != nil {
		return fmt.Errorf("failure to create server: %w", err)
	}

	if &server.RTMP != nil {
		server.RTMP.Start()
	}

	return nil
}
