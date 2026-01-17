#!/bin/bash

# Kitchen Sink - Request Inspection Test Script
BASE_URL="${BASE_URL:-http://localhost:8080}"

echo "=== Kitchen Sink - Request Inspection Tests ==="
echo "Base URL: $BASE_URL"
echo ""

# Test 1: Get IP
echo "1. Get IP:"
curl -s "$BASE_URL/api/v1/kitchen-sink/ip" | jq '.data'
echo ""

# Test 2: Get User-Agent
echo "2. Get User-Agent:"
curl -s -H "User-Agent: MyCustomCurlScript/1.0" "$BASE_URL/api/v1/kitchen-sink/user-agent" | jq '.data'
echo ""

# Test 3: Get Headers
echo "3. Get Headers:"
curl -s -H "X-My-Header: HelloWorld" "$BASE_URL/api/v1/kitchen-sink/headers" | jq '.data.headers["X-My-Header"]'
echo ""

echo "=== All tests completed ==="
