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
	// TODO: include rtmp.Server here

	// TODO: we may want to integrate rtmp.Server into server.Server
	server, err := server.NewServer(stream.NewPipeline(), detection.NewDetector())
	if err != nil {
		return fmt.Errorf("failure to create server: %w", err)
	}

	server.Run()

	return nil
}
