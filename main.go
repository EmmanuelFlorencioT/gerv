package main

import (
	"log"
	"github.com/EmmanuelFlorencioT/gerv/internal/http"
)

func main() {
	server := http.HTTPServer{
		Host: "127.0.0.1",
		Port: "8080",
	}

	if err := server.Serve(); err != nil {
		log.Fatalf("Failed to serve HTTP server: %w", err)
	}
}
