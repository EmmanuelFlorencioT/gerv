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
	parsedRequest, rErr := ParseRequest(conn)
	if rErr != nil {
		fmt.Println(rErr)
		return
	}

	resp := HandleRequest(*parsedRequest)

	// respond
	wErr := SendResponse(conn, *resp)
	if wErr != nil {
		fmt.Println(wErr)
		return
	}
}

func (server HTTPServer) Serve() {
	tcpServer := tcp.TCPServer{
		Host: server.Host,
		Port: server.Port,
	}

	tcpServer.ServeForever(server)
}
