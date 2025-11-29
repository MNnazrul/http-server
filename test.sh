#!/bin/bash

echo "=== HTTP Server Test Suite ==="
echo ""

SERVER_URL="http://localhost:4221"

echo "1. Testing root endpoint (/)..."
response=$(curl -s "$SERVER_URL/")
if [[ "$response" == *"Welcome to Go Server!"* ]]; then
    echo "   ✓ PASS: Root endpoint"
else
    echo "   ✗ FAIL: Root endpoint - Got: $response"
fi
echo ""

echo "2. Testing echo endpoint (/echo/hello)..."
response=$(curl -s "$SERVER_URL/echo/hello")
if [[ "$response" == "hello" ]]; then
    echo "   ✓ PASS: Echo endpoint"
else
    echo "   ✗ FAIL: Echo endpoint - Got: $response"
fi
echo ""

echo "3. Testing user-agent endpoint..."
response=$(curl -s "$SERVER_URL/user-agent")
if [[ -n "$response" ]]; then
    echo "   ✓ PASS: User-Agent endpoint - Got: $response"
else
    echo "   ✗ FAIL: User-Agent endpoint"
fi
echo ""

echo "4. Testing gzip compression..."
response=$(curl -s -H "Accept-Encoding: gzip" --compressed "$SERVER_URL/echo/test")
if [[ "$response" == "test" ]]; then
    echo "   ✓ PASS: Gzip compression"
else
    echo "   ✗ FAIL: Gzip compression - Got: $response"
fi
echo ""

echo "5. Testing 404 for unknown endpoint..."
status=$(curl -s -o /dev/null -w "%{http_code}" "$SERVER_URL/unknown")
if [[ "$status" == "404" ]]; then
    echo "   ✓ PASS: 404 handling"
else
    echo "   ✗ FAIL: 404 handling - Got status: $status"
fi
echo ""

echo "=== Test Suite Complete ==="

