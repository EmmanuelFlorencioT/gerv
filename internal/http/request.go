package http

import (
	"bufio"
	"errors"
	"net"
	"strings"
)

type HTTPRequest struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
}

func ParseRequest(conn net.Conn) (*HTTPRequest, error) {
	reader := bufio.NewReader(conn)

	// Parse request line
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, errors.New("could not read request line")
	}
	// clean up
	requestLine = strings.TrimSpace(requestLine)
	// split line into tokens
	tokens := strings.Fields(requestLine)
	if len(tokens) != 3 {
		return nil, errors.New("malformed http request line")
	}

	// read & store headers
	headers := make(map[string]string)
	for {
		headerLine, err := reader.ReadString('\n')
		if err != nil {
			return nil, errors.New("could not read headers")
		}
		if headerLine == "\r\n" || headerLine == "\n" {
			break
		}
		headerTokens := strings.Fields(headerLine)
		headers[headerTokens[0]] = headerTokens[1]
		// for educational purposes duplicate headers
		// will be overwritten
	}

	request := HTTPRequest{
		Method:  tokens[0],
		Path:    tokens[1],
		Version: tokens[2],
		Headers: headers,
	}

	return &request, nil
}
