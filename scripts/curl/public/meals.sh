#!/bin/bash

# Meals API Test Script
BASE_URL="${BASE_URL:-http://localhost:8080}"

echo "=== Meals API Tests ==="
echo "Base URL: $BASE_URL"
echo ""

# Test 1: Get all meals (default pagination)
echo "1. Get all meals (default pagination - page 1, limit 10):"
curl -s -X GET "$BASE_URL/api/v1/public/meals" | jq '.'
echo -e "\n"

# Test 2: Get meals with custom pagination
echo "2. Get meals with custom pagination (page 2, limit 5):"
curl -s -X GET "$BASE_URL/api/v1/public/meals?page=2&limit=5" | jq '.'
echo -e "\n"

# Test 3: Get meal by ID (valid)
echo "3. Get meal by ID (ID: 1):"
curl -s -X GET "$BASE_URL/api/v1/public/meals/1" | jq '.'
echo -e "\n"

# Test 4: Get meal by ID (valid - another meal)
echo "4. Get meal by ID (ID: 50):"
curl -s -X GET "$BASE_URL/api/v1/public/meals/50" | jq '.'
echo -e "\n"

# Test 5: Get meal by ID (invalid - not found)
echo "5. Get meal by ID (ID: 99999 - should fail):"
curl -s -X GET "$BASE_URL/api/v1/public/meals/99999" | jq '.'
echo -e "\n"

# Test 6: Get meal by ID (invalid - bad format)
echo "6. Get meal by ID (invalid format - should fail):"
curl -s -X GET "$BASE_URL/api/v1/public/meals/abc" | jq '.'
echo -e "\n"

# Test 7: Get random meal (call multiple times to verify randomness)
echo "7. Get a random meal (first call):"
curl -s -X GET "$BASE_URL/api/v1/public/meals/random" | jq '.data.id'
echo -e "\n"

echo "8. Get a random meal (second call):"
curl -s -X GET "$BASE_URL/api/v1/public/meals/random" | jq '.data.id'
echo -e "\n"

echo "9. Get a random meal (third call):"
curl -s -X GET "$BASE_URL/api/v1/public/meals/random" | jq '.data.id'
echo -e "\n"

# Test 10: Test pagination edge cases
echo "10. Test pagination - page beyond available data:"
curl -s -X GET "$BASE_URL/api/v1/public/meals?page=10000&limit=10" | jq '.'
echo -e "\n"

echo "=== All tests completed ==="
