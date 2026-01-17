#!/bin/bash

# Kitchen Sink - Status Codes Test Script
BASE_URL="${BASE_URL:-http://localhost:8080}"

echo "=== Kitchen Sink - Status Codes Tests ==="
echo "Base URL: $BASE_URL"
echo ""

# Test 1: 200 OK
echo "1. Status 200 (OK):"
curl -s -i "$BASE_URL/api/v1/kitchen-sink/status/200" | grep "HTTP"
echo ""

# Test 2: 201 Created
echo "2. Status 201 (Created):"
curl -s -i "$BASE_URL/api/v1/kitchen-sink/status/201" | grep "HTTP"
echo ""

# Test 3: 404 Not Found
echo "3. Status 404 (Not Found):"
curl -s -i "$BASE_URL/api/v1/kitchen-sink/status/404" | grep "HTTP"
echo ""

# Test 4: 418 I'm a teapot
echo "4. Status 418 (Teapot):"
curl -s -i "$BASE_URL/api/v1/kitchen-sink/status/418" | grep "HTTP"
echo ""

# Test 5: 500 Internal Server Error
echo "5. Status 500 (Internal Server Error):"
curl -s -i "$BASE_URL/api/v1/kitchen-sink/status/500" | grep "HTTP"
echo ""

echo "=== All tests completed ==="
