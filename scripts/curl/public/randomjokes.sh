#!/bin/bash

# Random Jokes API Test Script
BASE_URL="${BASE_URL:-http://localhost:8080}"

echo "=== Random Jokes API Tests ==="
echo "Base URL: $BASE_URL"
echo ""

# Test 1: Get all jokes (default pagination)
echo "1. Get all jokes (default pagination - page 1, limit 10):"
curl -s -X GET "$BASE_URL/api/v1/public/randomjokes" | jq '.'
echo -e "\n"

# Test 2: Get jokes with custom pagination
echo "2. Get jokes with custom pagination (page 2, limit 5):"
curl -s -X GET "$BASE_URL/api/v1/public/randomjokes?page=2&limit=5" | jq '.'
echo -e "\n"

# Test 3: Get joke by ID (valid)
echo "3. Get joke by ID (ID: 1):"
curl -s -X GET "$BASE_URL/api/v1/public/randomjokes/1" | jq '.'
echo -e "\n"

# Test 4: Get joke by ID (valid - another joke)
echo "4. Get joke by ID (ID: 50):"
curl -s -X GET "$BASE_URL/api/v1/public/randomjokes/50" | jq '.'
echo -e "\n"

# Test 5: Get joke by ID (invalid - not found)
echo "5. Get joke by ID (ID: 99999 - should fail):"
curl -s -X GET "$BASE_URL/api/v1/public/randomjokes/99999" | jq '.'
echo -e "\n"

# Test 6: Get joke by ID (invalid - bad format)
echo "6. Get joke by ID (invalid format - should fail):"
curl -s -X GET "$BASE_URL/api/v1/public/randomjokes/abc" | jq '.'
echo -e "\n"

# Test 7: Get random joke (call multiple times to verify randomness)
echo "7. Get a random joke (first call):"
curl -s -X GET "$BASE_URL/api/v1/public/randomjokes/random" | jq '.data.id'
echo -e "\n"

echo "8. Get a random joke (second call):"
curl -s -X GET "$BASE_URL/api/v1/public/randomjokes/random" | jq '.data.id'
echo -e "\n"

echo "9. Get a random joke (third call):"
curl -s -X GET "$BASE_URL/api/v1/public/randomjokes/random" | jq '.data.id'
echo -e "\n"

# Test 10: Test pagination edge cases
echo "10. Test pagination - page beyond available data:"
curl -s -X GET "$BASE_URL/api/v1/public/randomjokes?page=1000&limit=10" | jq '.'
echo -e "\n"

echo "=== All tests completed ==="
