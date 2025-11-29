package http

import (
	"strings"
)

type Request struct {
	Method  string
	Path    string
	Headers map[string]string
	Body    string
}

func ParseRequest(data string) *Request {
	lines := strings.Split(data, "\r\n")
	if len(lines) < 1 {
		return nil
	}

	parts := strings.Split(lines[0], " ")
	if len(parts) < 2 {
		return nil
	}

	method := parts[0]
	path := parts[1]
	headers := parseHeaders(lines)
	
	bodyParts := strings.SplitN(data, "\r\n\r\n", 2)
	body := ""
	if len(bodyParts) == 2 {
		body = bodyParts[1]
	}

	return &Request{
		Method:  method,
		Path:    path,
		Headers: headers,
		Body:    body,
	}
}

func parseHeaders(lines []string) map[string]string {
	headers := make(map[string]string)
	for _, line := range lines[1:] {
		if line == "" {
			break
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		headers[key] = value
	}
	return headers
}

