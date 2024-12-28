package stream

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lpipeline
#include "libpipeline/pipeline.h"
*/
import "C"

type Pipeline struct {
}

func NewPipeline() Pipeline {

	C.InitializeMediaPipeline()

	return Pipeline{}
}
