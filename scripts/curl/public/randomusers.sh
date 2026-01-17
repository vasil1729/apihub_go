#!/bin/bash

# Random Users API Test Script
BASE_URL="${BASE_URL:-http://localhost:8080}"

echo "=== Random Users API Tests ==="
echo "Base URL: $BASE_URL"
echo ""

# Test 1: Get all users (default pagination)
echo "1. Get all users (default pagination - page 1, limit 10):"
curl -s -X GET "$BASE_URL/api/v1/public/randomusers" | jq '.'
echo -e "\n"

# Test 2: Get users with custom pagination
echo "2. Get users with custom pagination (page 2, limit 5):"
curl -s -X GET "$BASE_URL/api/v1/public/randomusers?page=2&limit=5" | jq '.'
echo -e "\n"

# Test 3: Get user by ID (valid)
echo "3. Get user by ID (ID: 1):"
curl -s -X GET "$BASE_URL/api/v1/public/randomusers/1" | jq '.'
echo -e "\n"

# Test 4: Get user by ID (valid - another user)
echo "4. Get user by ID (ID: 5):"
curl -s -X GET "$BASE_URL/api/v1/public/randomusers/5" | jq '.'
echo -e "\n"

# Test 5: Get user by ID (invalid - not found)
echo "5. Get user by ID (ID: 99999 - should fail):"
curl -s -X GET "$BASE_URL/api/v1/public/randomusers/99999" | jq '.'
echo -e "\n"

# Test 6: Get user by ID (invalid - bad format)
echo "6. Get user by ID (invalid format - should fail):"
curl -s -X GET "$BASE_URL/api/v1/public/randomusers/abc" | jq '.'
echo -e "\n"

# Test 7: Get random user
echo "7. Get a random user:"
curl -s -X GET "$BASE_URL/api/v1/public/randomusers/random" | jq '.'
echo -e "\n"

# Test 8: Test pagination edge cases
echo "8. Test pagination - page beyond available data:"
curl -s -X GET "$BASE_URL/api/v1/public/randomusers?page=1000&limit=10" | jq '.'
echo -e "\n"

# Test 9: Test pagination - large limit
echo "9. Test pagination - large limit (should cap at 100):"
curl -s -X GET "$BASE_URL/api/v1/public/randomusers?page=1&limit=200" | jq '.'
echo -e "\n"

# Test 10: Test pagination - invalid parameters
echo "10. Test pagination - invalid page number (should default to 1):"
curl -s -X GET "$BASE_URL/api/v1/public/randomusers?page=-1&limit=10" | jq '.'
echo -e "\n"

echo "=== All tests completed ==="
