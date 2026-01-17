#!/bin/bash

# Kitchen Sink - HTTP Methods Test Script
BASE_URL="${BASE_URL:-http://localhost:8080}"

echo "=== Kitchen Sink - HTTP Methods Tests ==="
echo "Base URL: $BASE_URL"
echo ""

# Test 1: GET
echo "1. GET Request:"
curl -s -X GET "$BASE_URL/api/v1/kitchen-sink/http-methods/get?foo=bar&baz=qux" | jq '.data | {method, query, url}'
echo -e "\n"

# Test 2: POST with JSON
echo "2. POST Request (JSON):"
curl -s -X POST "$BASE_URL/api/v1/kitchen-sink/http-methods/post" \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "role": "Developer"}' | jq '.data | {method, body, headers}'
echo -e "\n"

# Test 3: PUT
echo "3. PUT Request:"
curl -s -X PUT "$BASE_URL/api/v1/kitchen-sink/http-methods/put" \
  -d "plain text data" | jq '.data | {method, body}'
echo -e "\n"

# Test 4: PATCH
echo "4. PATCH Request:"
curl -s -X PATCH "$BASE_URL/api/v1/kitchen-sink/http-methods/patch" | jq '.data.method'
echo -e "\n"

# Test 5: DELETE
echo "5. DELETE Request:"
curl -s -X DELETE "$BASE_URL/api/v1/kitchen-sink/http-methods/delete" | jq '.data.method'
echo -e "\n"

echo "=== All tests completed ==="
