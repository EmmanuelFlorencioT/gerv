package http

import (
	"errors"
	"net"
	"strconv"
	"strings"
)

type HTTPResponse struct {
	Protocol   string
	StatusCode int
	StatusText string
	Headers    map[string]string
	Data       string
}

func SendResponse(conn net.Conn, resp HTTPResponse) error {
	statusLine := resp.Protocol + " " + strconv.Itoa(resp.StatusCode) + " " + resp.StatusText + "\r\n"

	headers := make([]string, 0, len(resp.Headers))
	for key, value := range resp.Headers {
		headers = append(headers, key+": "+value)
	}

	response := statusLine +
		strings.Join(headers, "\r\n") +
		"\r\n\r\n" +
		resp.Data

	_, err := conn.Write([]byte(response))
	if err != nil {
		return errors.New("could not send response")
	}

	return nil
}
