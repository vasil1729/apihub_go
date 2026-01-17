#!/bin/bash

# Quotes API Test Script
BASE_URL="${BASE_URL:-http://localhost:8080}"

echo "=== Quotes API Tests ==="
echo "Base URL: $BASE_URL"
echo ""

# Test 1: Get all quotes (default pagination)
echo "1. Get all quotes (default pagination - page 1, limit 10):"
curl -s -X GET "$BASE_URL/api/v1/public/quotes" | jq '.'
echo -e "\n"

# Test 2: Get quotes with custom pagination
echo "2. Get quotes with custom pagination (page 2, limit 5):"
curl -s -X GET "$BASE_URL/api/v1/public/quotes?page=2&limit=5" | jq '.'
echo -e "\n"

# Test 3: Get quote by ID (valid)
echo "3. Get quote by ID (ID: 1):"
curl -s -X GET "$BASE_URL/api/v1/public/quotes/1" | jq '.'
echo -e "\n"

# Test 4: Get quote by ID (valid - another quote)
echo "4. Get quote by ID (ID: 50):"
curl -s -X GET "$BASE_URL/api/v1/public/quotes/50" | jq '.'
echo -e "\n"

# Test 5: Get quote by ID (invalid - not found)
echo "5. Get quote by ID (ID: 99999 - should fail):"
curl -s -X GET "$BASE_URL/api/v1/public/quotes/99999" | jq '.'
echo -e "\n"

# Test 6: Get quote by ID (invalid - bad format)
echo "6. Get quote by ID (invalid format - should fail):"
curl -s -X GET "$BASE_URL/api/v1/public/quotes/abc" | jq '.'
echo -e "\n"

# Test 7: Get random quote (call multiple times to verify randomness)
echo "7. Get a random quote (first call):"
curl -s -X GET "$BASE_URL/api/v1/public/quotes/random" | jq '.data.id'
echo -e "\n"

echo "8. Get a random quote (second call):"
curl -s -X GET "$BASE_URL/api/v1/public/quotes/random" | jq '.data.id'
echo -e "\n"

echo "9. Get a random quote (third call):"
curl -s -X GET "$BASE_URL/api/v1/public/quotes/random" | jq '.data.id'
echo -e "\n"

# Test 10: Test pagination edge cases
echo "10. Test pagination - page beyond available data:"
curl -s -X GET "$BASE_URL/api/v1/public/quotes?page=1000&limit=10" | jq '.'
echo -e "\n"

echo "=== All tests completed ==="
