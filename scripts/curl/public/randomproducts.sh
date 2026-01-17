#!/bin/bash

# Random Products API Test Script
BASE_URL="${BASE_URL:-http://localhost:8080}"

echo "=== Random Products API Tests ==="
echo "Base URL: $BASE_URL"
echo ""

# Test 1: Get all products (default pagination)
echo "1. Get all products (default pagination - page 1, limit 10):"
curl -s -X GET "$BASE_URL/api/v1/public/randomproducts" | jq '.pagination, .data | length'
echo -e "\n"

# Test 2: Get products with custom pagination
echo "2. Get products with custom pagination (page 2, limit 5):"
curl -s -X GET "$BASE_URL/api/v1/public/randomproducts?page=2&limit=5" | jq '.pagination'
echo -e "\n"

# Test 3: Get product by ID (valid - iPhone 9)
echo "3. Get product by ID (ID: 1 - iPhone 9):"
curl -s -X GET "$BASE_URL/api/v1/public/randomproducts/1" | jq '.data | {id, title, price, category}'
echo -e "\n"

# Test 4: Get product by ID (valid - another product)
echo "4. Get product by ID (ID: 10):"
curl -s -X GET "$BASE_URL/api/v1/public/randomproducts/10" | jq '.data | {id, title, price, category}'
echo -e "\n"

# Test 5: Get product by ID (invalid - not found)
echo "5. Get product by ID (ID: 99999 - should fail):"
curl -s -X GET "$BASE_URL/api/v1/public/randomproducts/99999" | jq '.'
echo -e "\n"

# Test 6: Get product by ID (invalid - bad format)
echo "6. Get product by ID (invalid format - should fail):"
curl -s -X GET "$BASE_URL/api/v1/public/randomproducts/abc" | jq '.'
echo -e "\n"

# Test 7: Get random product (call multiple times to verify randomness)
echo "7. Get a random product (first call):"
curl -s -X GET "$BASE_URL/api/v1/public/randomproducts/random" | jq '.data | {id, title, price}'
echo -e "\n"

echo "8. Get a random product (second call):"
curl -s -X GET "$BASE_URL/api/v1/public/randomproducts/random" | jq '.data | {id, title, price}'
echo -e "\n"

echo "9. Get a random product (third call):"
curl -s -X GET "$BASE_URL/api/v1/public/randomproducts/random" | jq '.data | {id, title, price}'
echo -e "\n"

# Test 10: Test pagination edge cases
echo "10. Test pagination - page beyond available data:"
curl -s -X GET "$BASE_URL/api/v1/public/randomproducts?page=10000&limit=10" | jq '.pagination, .data | length'
echo -e "\n"

# Test 11: Test large limit
echo "11. Test large limit (limit 50):"
curl -s -X GET "$BASE_URL/api/v1/public/randomproducts?page=1&limit=50" | jq '.pagination, .data | length'
echo -e "\n"

# Test 12: Filter by checking response structure
echo "12. Verify product structure (checking rating):"
curl -s -X GET "$BASE_URL/api/v1/public/randomproducts/1" | jq '.data | {id, title, price, rating}'
echo -e "\n"

echo "=== All tests completed ==="
