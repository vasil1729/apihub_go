#!/bin/bash

# Kitchen Sink - Redirects Test Script
BASE_URL="${BASE_URL:-http://localhost:8080}"

echo "=== Kitchen Sink - Redirects Tests ==="
echo "Base URL: $BASE_URL"
echo ""

# Test 1: 301
echo "1. 301 Moved Permanently:"
curl -s -i "$BASE_URL/api/v1/kitchen-sink/redirects/301" | grep "Location"
echo ""

# Test 2: 302
echo "2. 302 Found (Default to google.com):"
curl -s -i "$BASE_URL/api/v1/kitchen-sink/redirects/302" | grep "Location"
echo ""

# Test 3: 302 Custom URL
echo "3. 302 Found (Custom URL):"
curl -s -i "$BASE_URL/api/v1/kitchen-sink/redirects/302?url=https://github.com" | grep "Location"
echo ""

echo "=== All tests completed ==="
