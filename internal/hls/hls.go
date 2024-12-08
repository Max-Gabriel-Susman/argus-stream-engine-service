package hls

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Server struct {
	ContentDirectory string
}

func NewServer(contentDirectory string) Server {
	return Server{
		ContentDirectory: contentDirectory,
	}
}

func (s Server) ServeContent() {
	if _, err := os.Stat(s.ContentDirectory); os.IsNotExist(err) {
		log.Fatalf("HLS directory %s does not exist", s.ContentDirectory)
	}

	fs := http.FileServer(http.Dir(s.ContentDirectory))
	http.Handle("/hls/", http.StripPrefix("/hls/", fs))

	port := 8080
	fmt.Printf("Starting HLS server on http://localhost:%d/hls/\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
