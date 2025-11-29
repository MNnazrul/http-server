package http

import (
	"bytes"
	"compress/gzip"
	"fmt"
)

const (
	StatusOK             = "200 OK"
	StatusCreated        = "201 Created"
	StatusBadRequest     = "400 Bad Request"
	StatusNotFound       = "404 Not Found"
	StatusInternalError  = "500 Internal Server Error"
	ContentTypePlain     = "text/plain"
	ContentTypeOctet     = "application/octet-stream"
	ContentEncodingGzip  = "gzip"
	ConnectionClose      = "close"
)

type Response struct {
	StatusCode    string
	Headers       map[string]string
	Body          []byte
	Connection    string
}

func NewResponse(statusCode string) *Response {
	return &Response{
		StatusCode: statusCode,
		Headers:    make(map[string]string),
	}
}

func (r *Response) SetContentType(contentType string) {
	r.Headers["Content-Type"] = contentType
}

func (r *Response) SetBody(body []byte) {
	r.Body = body
	r.Headers["Content-Length"] = fmt.Sprintf("%d", len(body))
}

func (r *Response) SetConnection(connection string) {
	r.Connection = connection
}

func (r *Response) CompressGzip() error {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	if _, err := gz.Write(r.Body); err != nil {
		return err
	}
	if err := gz.Close(); err != nil {
		return err
	}
	r.Body = buf.Bytes()
	r.Headers["Content-Encoding"] = ContentEncodingGzip
	r.Headers["Content-Length"] = fmt.Sprintf("%d", len(r.Body))
	return nil
}

func (r *Response) Bytes() []byte {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("HTTP/1.1 %s\r\n", r.StatusCode))
	
	for key, value := range r.Headers {
		buf.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}
	
	if r.Connection != "" {
		buf.WriteString(fmt.Sprintf("Connection: %s\r\n", r.Connection))
	}
	
	buf.WriteString("\r\n")
	buf.Write(r.Body)
	
	return buf.Bytes()
}

