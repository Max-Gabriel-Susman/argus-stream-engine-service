package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lpipeline
#include "../../pipeline.h"
*/
import "C"
import "github.com/Max-Gabriel-Susman/argus-stream-engine-service/internal/detection"

func main() {
	pipeline := pipeline.NewPipeline()

	detector := detection.NewDetector()

	C.InitializeMediaPipeline()
}
