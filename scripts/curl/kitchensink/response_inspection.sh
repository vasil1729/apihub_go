#!/bin/bash

# Kitchen Sink - Response Inspection Test Script
BASE_URL="${BASE_URL:-http://localhost:8080}"

echo "=== Kitchen Sink - Response Inspection Tests ==="
echo "Base URL: $BASE_URL"
echo ""

# Test 1: JSON
echo "1. JSON Response:"
curl -s -i "$BASE_URL/api/v1/kitchen-sink/response/json" | head -n 15
echo ""
echo ""

# Test 2: XML
echo "2. XML Response:"
curl -s -i "$BASE_URL/api/v1/kitchen-sink/response/xml" | head -n 15
echo ""
echo ""

# Test 3: HTML
echo "3. HTML Response:"
curl -s -i "$BASE_URL/api/v1/kitchen-sink/response/html" | head -n 15
echo ""

echo "=== All tests completed ==="
