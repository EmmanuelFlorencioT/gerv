package http

func HandleRequest(req HTTPRequest) *HTTPResponse {
	if req.Method != "GET" {
		resp := HTTPResponse{
			Protocol:   "HTTP/1.1",
			StatusCode: 405,
			StatusText: "METHOD NOT ALLOWED",
			Headers:    make(map[string]string),
			Data:       "",
		}
		return &resp
	}

	if req.Path != "/" {
		resp := HTTPResponse{
			Protocol:   "HTTP/1.1",
			StatusCode: 404,
			StatusText: "NOT FOUND",
			Headers:    make(map[string]string),
			Data:       "",
		}
		return &resp
	}

	resp := HTTPResponse{
		Protocol:   "HTTP/1.1",
		StatusCode: 200,
		StatusText: "OK",
		Headers: map[string]string{
			"Content-Type":   "text/plain; charset=utf-8",
			"Content-Lenght": "14",
		},
		Data: "Route matched!",
	}

	return &resp
}
