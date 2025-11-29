# How to Run and Test the HTTP Server

## Prerequisites
- Go 1.20+ installed
- `go` command available in PATH

## Running the Server

### Basic Run (without directory)
```bash
go run cmd/server/main.go
```

### Run with Directory (for file serving)
```bash
go run cmd/server/main.go -directory /path/to/your/directory
```

### Build and Run
```bash
# Build the binary
go build -o http-server cmd/server/main.go

# Run the binary
./http-server

# Or with directory
./http-server -directory /path/to/your/directory
```

The server will start on **port 4221** by default.

## Testing the Server

### 1. Test Root Endpoint
```bash
curl http://localhost:4221/
```

Expected: `Welcome to Go Server!`

### 2. Test Echo Endpoint
```bash
curl http://localhost:4221/echo/hello
```

Expected: `hello`

### 3. Test Echo with Gzip Compression
```bash
curl -H "Accept-Encoding: gzip" --compressed http://localhost:4221/echo/hello
```

### 4. Test User-Agent Endpoint
```bash
curl http://localhost:4221/user-agent
```

Expected: Your curl user-agent string

### 5. Test File Serving (GET)
```bash
# First, create a test directory and file
mkdir -p /tmp/test-files
echo "Hello from file" > /tmp/test-files/test.txt

# Run server with directory
go run cmd/server/main.go -directory /tmp/test-files

# In another terminal, test file retrieval
curl http://localhost:4221/files/test.txt
```

### 6. Test File Upload (POST)
```bash
# Server should be running with -directory flag
curl -X POST http://localhost:4221/files/uploaded.txt \
  -d "This is the file content"
```

### 7. Test Connection: Close
```bash
curl -H "Connection: close" http://localhost:4221/
```

## Quick Test Script

Save this as `test.sh` and run: `chmod +x test.sh && ./test.sh`

```bash
#!/bin/bash
echo "Testing HTTP Server..."

# Test root
echo "1. Testing root endpoint..."
curl -s http://localhost:4221/ | grep -q "Welcome" && echo "✓ Root works" || echo "✗ Root failed"

# Test echo
echo "2. Testing echo endpoint..."
curl -s http://localhost:4221/echo/test123 | grep -q "test123" && echo "✓ Echo works" || echo "✗ Echo failed"

# Test user-agent
echo "3. Testing user-agent endpoint..."
curl -s http://localhost:4221/user-agent | grep -q "curl" && echo "✓ User-Agent works" || echo "✗ User-Agent failed"

echo "Tests complete!"
```

