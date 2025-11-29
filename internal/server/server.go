package server

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/codecrafters-io/http-server-starter-go/internal/handler"
	"github.com/codecrafters-io/http-server-starter-go/internal/http"
)

const (
	DefaultPort = "4221"
	BufferSize  = 1024
)

type Server struct {
	port      string
	handler   *handler.Handler
	listener  net.Listener
}

func NewServer(port string, directory string) *Server {
	return &Server{
		port:    port,
		handler: handler.NewHandler(directory),
	}
}

func (s *Server) Start() error {
	address := fmt.Sprintf("0.0.0.0:%s", s.port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to bind to port %s: %w", s.port, err)
	}
	
	s.listener = listener
	fmt.Printf("Listening on port %s...\n", s.port)
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	
	buffer := make([]byte, BufferSize)
	
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err != io.EOF {
				log.Printf("Read error: %v", err)
			}
			return
		}
		
		request := http.ParseRequest(string(buffer[:n]))
		if request == nil {
			continue
		}
		
		response := s.handler.HandleRequest(request)
		s.setConnectionHeader(request, response)
		
		if _, err := conn.Write(response.Bytes()); err != nil {
			log.Printf("Write error: %v", err)
			return
		}
		
		if response.Connection == http.ConnectionClose {
			return
		}
	}
}

func (s *Server) setConnectionHeader(req *http.Request, resp *http.Response) {
	if connection, exists := req.Headers["Connection"]; exists {
		if connection == http.ConnectionClose {
			resp.SetConnection(http.ConnectionClose)
		}
	}
}

func (s *Server) Close() error {
	if s.listener != nil {
		return s.listener.Close()
	}
	return nil
}

