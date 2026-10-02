package tcp

import (
	"fmt"
	"net"
)

type TCPServer struct {
	Host string
	Port string
}

type ConnectionHandler interface {
	HandleConnection(net.Conn)
}

func (server TCPServer) ServeForever(handler ConnectionHandler) {
	listener, err := net.Listen("tcp", server.Host+":"+server.Port)
	if err != nil {
		return
	}

	defer listener.Close()

	fmt.Println("Listening on port " + server.Port + "...")

	// Serve forever
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection", err)
			continue
		}

		// do something with
		// established connection
		go handler.HandleConnection(conn)
	}
}
