#!/bin/bash

# Stocks API Test Script
BASE_URL="${BASE_URL:-http://localhost:8080}"

echo "=== Stocks API Tests ==="
echo "Base URL: $BASE_URL"
echo ""

# Test 1: Get all stocks (default pagination)
echo "1. Get all stocks (default pagination - page 1, limit 10):"
curl -s -X GET "$BASE_URL/api/v1/public/stocks" | jq '.'
echo -e "\n"

# Test 2: Get stocks with custom pagination
echo "2. Get stocks with custom pagination (page 2, limit 5):"
curl -s -X GET "$BASE_URL/api/v1/public/stocks?page=2&limit=5" | jq '.'
echo -e "\n"

# Test 3: Get stock by symbol (TCS)
echo "3. Get stock by symbol (TCS):"
curl -s -X GET "$BASE_URL/api/v1/public/stocks/TCS" | jq '.'
echo -e "\n"

# Test 4: Get stock by symbol (RELIANCE)
echo "4. Get stock by symbol (RELIANCE):"
curl -s -X GET "$BASE_URL/api/v1/public/stocks/RELIANCE" | jq '.'
echo -e "\n"

# Test 5: Get stock by symbol (INFY)
echo "5. Get stock by symbol (INFY):"
curl -s -X GET "$BASE_URL/api/v1/public/stocks/INFY" | jq '.'
echo -e "\n"

# Test 6: Get stock by symbol - case insensitive (lowercase)
echo "6. Get stock by symbol - case insensitive (hdfc):"
curl -s -X GET "$BASE_URL/api/v1/public/stocks/hdfc" | jq '.'
echo -e "\n"

# Test 7: Get stock by symbol (invalid - should fail)
echo "7. Get stock by symbol (NOTEXIST - should fail):"
curl -s -X GET "$BASE_URL/api/v1/public/stocks/NOTEXIST" | jq '.'
echo -e "\n"

# Test 8: Get stock by symbol (empty - should fail)
echo "8. Get stock by symbol (empty - should fail):"
curl -s -X GET "$BASE_URL/api/v1/public/stocks/" | jq '.'
echo -e "\n"

# Test 9: Test pagination - large page number
echo "9. Test pagination - page 100, limit 20:"
curl -s -X GET "$BASE_URL/api/v1/public/stocks?page=100&limit=20" | jq '.data | length'
echo -e "\n"

# Test 10: Test pagination - page beyond available data
echo "10. Test pagination - page beyond available data:"
curl -s -X GET "$BASE_URL/api/v1/public/stocks?page=10000&limit=10" | jq '.'
echo -e "\n"

echo "=== All tests completed ==="
