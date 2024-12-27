package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lpipeline
#include "../../pipeline.h"
*/
import "C"
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
	C.InitializeMediaPipeline()

	server, err := server.NewServer(stream.NewPipeline(), detection.NewDetector())
	if err != nil {
		return fmt.Errorf("failure to create server: %w", err)
	}

	server.Run()

	return nil
}
