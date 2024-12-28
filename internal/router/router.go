package router

import (
	"github.com/Max-Gabriel-Susman/argus-stream-engine-service/internal/detection"
	"github.com/Max-Gabriel-Susman/argus-stream-engine-service/internal/stream"
)

type Router struct {
	Pipeline stream.Pipeline
	Detector detection.Detector
}

func NewRouter(pipeline stream.Pipeline, detector detection.Detector) Router {
	return Router{
		Pipeline: pipeline,
		Detector: detector,
	}
}