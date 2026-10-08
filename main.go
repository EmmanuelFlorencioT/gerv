package main

import (
	"github.com/EmmanuelFlorencioT/gerv/internal/http"
)

func main() {
	server := http.HTTPServer{
		Host: "127.0.0.1",
		Port: "8080",
	}

	server.Serve()
}
