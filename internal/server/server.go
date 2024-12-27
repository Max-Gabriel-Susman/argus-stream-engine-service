package server

import (
	"log"

	"github.com/Max-Gabriel-Susman/argus-stream-engine-service/internal/detection"
	"github.com/Max-Gabriel-Susman/argus-stream-engine-service/internal/router"
	"github.com/Max-Gabriel-Susman/argus-stream-engine-service/internal/stream"
)

type Server struct {
	router.Router
}

func NewServer(pipeline stream.Pipeline, detector detection.Detector) (Server, error) {

	router := router.NewRouter(pipeline, detector)

	return Server{
		Router: router,
	}, nil
}

func (s Server) Run() {
	log.Println("Initializing Argus Stream Engine Service...")

	log.Println("Argus Stream Engine Service online.")
}
