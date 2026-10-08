package http

import (
	"fmt"
	"net"

	"github.com/EmmanuelFlorencioT/gerv/internal/tcp"
)

type HTTPServer struct {
	Host string
	Port string
}

func (server HTTPServer) HandleConnection(conn net.Conn) {
	defer conn.Close()
	// read bytes & parse request
	parsedRequest, err := ParseRequest(conn)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp := HandleRequest(*parsedRequest)

	// respond
	// respond with desired resource if available
	// respond with failure if not available
	// close connection
}

func (server HTTPServer) Serve() {
	tcpServer := tcp.TCPServer{
		Host: server.Host,
		Port: server.Port,
	}

	tcpServer.ServeForever(server)
}
