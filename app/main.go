package main

import (
	"bytes"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

var dir string

func main() {
	dir = inputDirCommand()

	listener, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Listening on port 4221...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err != io.EOF {
				log.Println("Read error:", err)
			}
			return
		}

		request := string(buffer[:n])
		lines := strings.Split(request, "\r\n")
		if len(lines) < 1 {
			continue
		}

		part := strings.Split(lines[0], " ")
		if len(part) < 2 {
			continue
		}

		method := part[0]
		path := part[1]
		headers := parseHeaders(lines)
		var res string

		fmt.Println(headers)

		if method == "GET" {
			if path == "/user-agent" {
				str, ok := headers["User-Agent"]
				if !ok {
					res = "HTTP/1.1 404 Not Found\r\n\r\n"
				} else {
					res = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(str), str)
				}
			} else if path == "/" {
				body := "Welcome to Go Server!"
				res = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(body), body)
			} else if strings.HasPrefix(path, "/echo/") {
				str1 := strings.TrimPrefix(path, "/echo/")
				headersList := strings.Split(headers["Accept-Encoding"], ",")

				useGzip := false
				for _, acceptEncoding := range headersList {
					str := strings.TrimSpace(acceptEncoding)
					if str == "gzip" {
						useGzip = true
						break
					}
				}

				debugHex := false
				if val, ok := headers["X-Debug-Hex"]; ok && strings.TrimSpace(val) == "true" {
					debugHex = true
				}

				if useGzip {
					var buf bytes.Buffer
					gz := gzip.NewWriter(&buf)
					_, err := gz.Write([]byte(str1))
					if err != nil {
						panic(err)
					}
					gz.Close()
					if debugHex {
						hexStr := fmt.Sprintf("%x", buf.Bytes())
						res = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Encoding: gzip\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%x", len(hexStr), hexStr)
					} else {
						res = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Encoding: gzip\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(buf.Bytes()), buf.Bytes())
					}
				} else {
					res = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(str1), str1)
				}
			} else if strings.HasPrefix(path, "/files/") {
				// fmt.Println("inside the files route")
				filename := strings.TrimPrefix(path, "/files/")
				filePath := fmt.Sprintf("%s/%s", strings.TrimRight(dir, "/"), filename)
				// fmt.Println(filename)
				// fmt.Println(filePath)
				if fileExistts(filePath) {
					content, err := os.ReadFile(filePath)
					if err != nil {
						res = "HTTP/1.1 404 Not Found\r\n\r\n"
					} else {
						res = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: application/octet-stream\r\nContent-Length: %d\r\n\r\n%s", len(content), content)
					}
				} else {
					res = "HTTP/1.1 404 Not Found\r\n\r\n"
				}
			} else {
				res = "HTTP/1.1 404 Not Found\r\n\r\n"
			}
		} else if method == "POST" {
			filename := strings.TrimPrefix(path, "/files/")
			filePath := fmt.Sprintf("%s/%s", strings.TrimRight(dir, "/"), filename)

			parts := strings.SplitN(request, "\r\n\r\n", 2)
			// fmt.Println(parts)
			if len(parts) != 2 {
				res = "HTTP/1.1 400 Bad Request\r\n\r\n"
			} else {
				body := parts[1]
				// fmt.Println(body)
				// Write to file
				err := os.WriteFile(filePath, []byte(body), 0o644)
				if err != nil {
					res = "HTTP/1.1 500 Internal Server Error\r\n\r\n"
				} else {
					res = "HTTP/1.1 201 Created\r\n\r\n"
				}
			}
		} else {
			res = "HTTP/1.1 404 Not Found\r\n\r\n"
		}

		closeConnection := false
		if _, ok := headers["Connection"]; ok {
			if headers["Connection"] == "close" {
				closeConnection = true
				res = addConnectionHeader(res)
			}
		}
		conn.Write([]byte(res))
		if closeConnection {
			os.Exit(0)
		}
	}
}

func addConnectionHeader(res string) string {
	parts := strings.SplitN(res, "\r\n\r\n", 2)
	if len(parts) != 2 {
		return res
	}

	headers := parts[0]
	body := parts[1]

	lines := strings.Split(headers, "\r\n")
	if len(lines) == 0 {
		return res
	}

	statusLine := lines[0]
	if strings.Contains(statusLine, "200") || strings.Contains(statusLine, "201") {
		// Insert Connection: close after status line
		newHeaders := []string{statusLine, "Connection: close"}
		newHeaders = append(newHeaders, lines[1:]...)
		newHeaderStr := strings.Join(newHeaders, "\r\n")
		return newHeaderStr + "\r\n\r\n" + body
	}

	return res
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

func fileExistts(path string) bool {
	_, err := os.Stat(path)
	return err != nil || !os.IsNotExist(err)
}

func inputDirCommand() string {
	dirPtr := flag.String("directory", "", "Path to directory")
	flag.Parse()
	fmt.Println(*dirPtr)
	return *dirPtr
}
