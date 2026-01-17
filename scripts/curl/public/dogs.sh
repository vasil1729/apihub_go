#!/bin/bash

# Dogs API Test Script
BASE_URL="${BASE_URL:-http://localhost:8080}"

echo "=== Dogs API Tests ==="
echo "Base URL: $BASE_URL"
echo ""

echo "1. Get all dog breeds (default pagination):"
curl -s -X GET "$BASE_URL/api/v1/public/dogs" | jq '.'
echo -e "\n"

echo "2. Get dog breeds with custom pagination (page 2, limit 5):"
curl -s -X GET "$BASE_URL/api/v1/public/dogs?page=2&limit=5" | jq '.'
echo -e "\n"

echo "3. Get dog breed by ID (ID: 1 - Affenpinscher):"
curl -s -X GET "$BASE_URL/api/v1/public/dogs/1" | jq '.'
echo -e "\n"

echo "4. Get dog breed by ID (ID: 29 - Beagle):"
curl -s -X GET "$BASE_URL/api/v1/public/dogs/29" | jq '.'
echo -e "\n"

echo "5. Get dog breed by ID (invalid - should fail):"
curl -s -X GET "$BASE_URL/api/v1/public/dogs/99999" | jq '.'
echo -e "\n"

echo "6. Get random dog breed (first call):"
curl -s -X GET "$BASE_URL/api/v1/public/dogs/random" | jq '.data.name'
echo -e "\n"

echo "7. Get random dog breed (second call):"
curl -s -X GET "$BASE_URL/api/v1/public/dogs/random" | jq '.data.name'
echo -e "\n"

echo "8. Get random dog breed (third call):"
curl -s -X GET "$BASE_URL/api/v1/public/dogs/random" | jq '.data.name'
echo -e "\n"

echo "=== All tests completed ==="
