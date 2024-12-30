package server

import (
	"log"

	"github.com/Max-Gabriel-Susman/rtmp"

	"github.com/Max-Gabriel-Susman/argus-stream-engine-service/internal/router"
	"github.com/Max-Gabriel-Susman/argus-stream-engine-service/internal/stream"
)

type Server struct {
	Router router.Router
	RTMP   rtmp.Server
}

func NewStreamServer(pipeline stream.Pipeline) (Server, error) {
	router := router.NewRouter(pipeline)

	return Server{
		Router: router,
		RTMP:   *rtmp.NewServer(),
	}, nil
}

func (s Server) Run() {
	log.Println("Initializing Argus Stream Engine Service...")

	log.Println("Argus Stream Engine Service online.")
}
