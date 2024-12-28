package detection

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -ldetector
#include "libdetector/detector.h"
*/
import "C"

type Detector struct {
}

func NewDetector() Detector {
	C.InitializeObjectDetection()

	return Detector{}
}
