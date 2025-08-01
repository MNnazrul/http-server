Here's a `README.md` you can use for your GitHub repository, describing your implementation of the **Build Your Own HTTP Server** challenge in Go:

---

# 🕸️ Build Your Own HTTP Server in Go

This is my implementation of the [CodeCrafters](https://codecrafters.io) "Build Your Own HTTP Server" challenge using the Go programming language.

> 🚀 Challenge Completed: [View Progress](https://app.codecrafters.io/users/codecrafters-bot?r=2qF)

## 📖 About the Challenge

In this project, I built a simple yet functional HTTP/1.1 server from scratch using raw TCP sockets and the Go standard library.

Through this, I learned:

- How HTTP/1.1 requests are structured
- How to parse HTTP methods, headers, and bodies
- How to respond with correct HTTP responses
- How to handle persistent connections (`Connection: keep-alive` and `Connection: close`)
- How to support gzip-encoded responses
- How to serve files from disk and echo request paths

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
chmod +x your_program.sh
./your_program.sh
```

By default, the server listens on port `4221`. You can test it with:

```bash
curl -v http://localhost:4221/echo/hello
curl -v -H "Accept-Encoding: gzip" http://localhost:4221/echo/hello
```

## 📂 Project Structure

```
.
├── app
│   └── main.go        # Main server logic
├── your_program.sh    # Entrypoint script used by CodeCrafters
├── README.md
```

## 🧩 Based On

This project is based on the challenge from [CodeCrafters.io](https://codecrafters.io/courses/http-server/overview), where you build projects by replicating systems from scratch.

## 📜 License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
