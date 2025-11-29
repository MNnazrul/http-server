Here's a `README.md` you can use for your GitHub repository, describing your implementation of the **Build Your Own HTTP Server** challenge in Go:

---

# 🕸️ Build Your Own HTTP Server in Go

This is my implementation of the [CodeCrafters](https://codecrafters.io) "Build Your Own HTTP Server" challenge using the Go programming language.

> 🚀 Challenge Completed: [View Progress](https://app.codecrafters.io/users/codecrafters-bot?r=2qF)

## 📖 About the Challenge

In this project, I built a simple yet functional HTTP/1.1 server from scratch using raw TCP sockets and the Go standard library. The codebase has been refactored to follow clean architecture principles with proper separation of concerns.

Through this, I learned:

- How HTTP/1.1 requests are structured
- How to parse HTTP methods, headers, and bodies
- How to respond with correct HTTP responses
- How to handle persistent connections (`Connection: keep-alive` and `Connection: close`)
- How to support gzip-encoded responses
- How to serve files from disk and echo request paths
- Clean code architecture and separation of concerns

## ✅ Completed Features

- [x] TCP server on a custom port
- [x] Support for `GET` requests
- [x] `/echo/:param` path that returns the value in the URL
- [x] Add custom response headers
- [x] Support `Connection: close` and `keep-alive`
- [x] Support for `gzip` compression if `Accept-Encoding: gzip` is present
- [x] Serve static files from a directory
- [x] Correctly handle multiple clients

## 🧠 Key Learnings

- Low-level networking with Go's `net` package
- Manually parsing and writing HTTP requests/responses
- Understanding content encoding (gzip)
- Working with concurrency and sockets

## 🛠️ How to Run Locally

Make sure you have Go 1.20+ installed.

```bash
git clone https://github.com/<your-username>/http-server-go.git
cd http-server-go

# Run the server
go run cmd/server/main.go

# Or with directory for file serving
go run cmd/server/main.go -directory /path/to/directory

# Build and run
go build -o http-server cmd/server/main.go
./http-server -directory /path/to/directory
```

By default, the server listens on port `4221`. You can test it with:

```bash
# Basic endpoints
curl http://localhost:4221/
curl http://localhost:4221/echo/hello
curl http://localhost:4221/user-agent

# With gzip compression
curl -H "Accept-Encoding: gzip" --compressed http://localhost:4221/echo/hello

# File operations (requires -directory flag)
curl http://localhost:4221/files/test.txt
curl -X POST http://localhost:4221/files/upload.txt -d "content"
```

Run the automated test suite:
```bash
./test.sh
```

## 📂 Project Structure

```
.
├── cmd/
│   └── server/
│       └── main.go           # Application entry point
├── internal/
│   ├── server/
│   │   └── server.go         # TCP server and connection handling
│   ├── handler/
│   │   └── handler.go        # HTTP request handlers (routes)
│   └── http/
│       ├── request.go         # HTTP request parsing
│       └── response.go       # HTTP response building
├── go.mod
├── go.sum
├── README.md
├── RUN.md                     # Detailed run instructions
└── test.sh                    # Automated test suite
```

## 🏗️ Architecture

The project follows clean architecture principles with clear separation of concerns:

- **`cmd/server`**: Application entry point and CLI flag parsing
- **`internal/server`**: TCP server setup, connection management, and request routing
- **`internal/handler`**: Business logic for handling different HTTP endpoints
- **`internal/http`**: HTTP protocol implementation (request parsing and response building)

This structure ensures:
- ✅ Single responsibility principle
- ✅ Easy testing and maintenance
- ✅ Clear separation between networking, business logic, and protocol handling
- ✅ No global state or shared mutable variables

## 🧩 Based On

This project is based on the challenge from [CodeCrafters.io](https://codecrafters.io/courses/http-server/overview), where you build projects by replicating systems from scratch.

## 📜 License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
