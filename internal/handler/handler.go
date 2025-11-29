package handler

import (
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/internal/http"
)

type Handler struct {
	directory string
}

func NewHandler(directory string) *Handler {
	return &Handler{directory: directory}
}

func (h *Handler) HandleRequest(req *http.Request) *http.Response {
	if req.Method == "GET" {
		return h.handleGET(req)
	}
	if req.Method == "POST" {
		return h.handlePOST(req)
	}
	return http.NewResponse(http.StatusNotFound)
}

func (h *Handler) handleGET(req *http.Request) *http.Response {
	if req.Path == "/" {
		return h.handleRoot()
	}
	if req.Path == "/user-agent" {
		return h.handleUserAgent(req)
	}
	if strings.HasPrefix(req.Path, "/echo/") {
		return h.handleEcho(req)
	}
	if strings.HasPrefix(req.Path, "/files/") {
		return h.handleGetFile(req)
	}
	return http.NewResponse(http.StatusNotFound)
}

func (h *Handler) handleRoot() *http.Response {
	body := "Welcome to Go Server!"
	resp := http.NewResponse(http.StatusOK)
	resp.SetContentType(http.ContentTypePlain)
	resp.SetBody([]byte(body))
	return resp
}

func (h *Handler) handleUserAgent(req *http.Request) *http.Response {
	userAgent, exists := req.Headers["User-Agent"]
	if !exists {
		return http.NewResponse(http.StatusNotFound)
	}
	resp := http.NewResponse(http.StatusOK)
	resp.SetContentType(http.ContentTypePlain)
	resp.SetBody([]byte(userAgent))
	return resp
}

func (h *Handler) handleEcho(req *http.Request) *http.Response {
	message := strings.TrimPrefix(req.Path, "/echo/")
	resp := http.NewResponse(http.StatusOK)
	resp.SetContentType(http.ContentTypePlain)
	resp.SetBody([]byte(message))
	
	if h.shouldCompress(req) {
		if err := resp.CompressGzip(); err != nil {
			return http.NewResponse(http.StatusInternalError)
		}
	}
	
	return resp
}

func (h *Handler) handleGetFile(req *http.Request) *http.Response {
	filename := strings.TrimPrefix(req.Path, "/files/")
	filePath := h.buildFilePath(filename)
	
	if !fileExists(filePath) {
		return http.NewResponse(http.StatusNotFound)
	}
	
	content, err := os.ReadFile(filePath)
	if err != nil {
		return http.NewResponse(http.StatusNotFound)
	}
	
	resp := http.NewResponse(http.StatusOK)
	resp.SetContentType(http.ContentTypeOctet)
	resp.SetBody(content)
	return resp
}

func (h *Handler) handlePOST(req *http.Request) *http.Response {
	if !strings.HasPrefix(req.Path, "/files/") {
		return http.NewResponse(http.StatusNotFound)
	}
	
	filename := strings.TrimPrefix(req.Path, "/files/")
	filePath := h.buildFilePath(filename)
	
	if err := os.WriteFile(filePath, []byte(req.Body), 0644); err != nil {
		return http.NewResponse(http.StatusInternalError)
	}
	
	return http.NewResponse(http.StatusCreated)
}

func (h *Handler) shouldCompress(req *http.Request) bool {
	acceptEncoding, exists := req.Headers["Accept-Encoding"]
	if !exists {
		return false
	}
	
	encodings := strings.Split(acceptEncoding, ",")
	for _, encoding := range encodings {
		if strings.TrimSpace(encoding) == http.ContentEncodingGzip {
			return true
		}
	}
	return false
}

func (h *Handler) buildFilePath(filename string) string {
	dir := strings.TrimRight(h.directory, "/")
	return fmt.Sprintf("%s/%s", dir, filename)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

