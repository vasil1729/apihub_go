#!/bin/bash

# YouTube API Test Script
BASE_URL="${BASE_URL:-http://localhost:8080}"

echo "=== YouTube API Tests ==="
echo "Base URL: $BASE_URL"
echo ""

# Test 1: Get all videos (default pagination)
echo "1. Get all videos (default pagination - page 1, limit 10):"
curl -s -X GET "$BASE_URL/api/v1/public/youtube" | jq '.pagination, .data | length'
echo -e "\n"

# Test 2: Get videos with custom pagination
echo "2. Get videos with custom pagination (page 2, limit 5):"
curl -s -X GET "$BASE_URL/api/v1/public/youtube?page=2&limit=5" | jq '.pagination'
echo -e "\n"

# Test 3: Get video by ID (valid)
echo "3. Get video by ID (ID: 1):"
curl -s -X GET "$BASE_URL/api/v1/public/youtube/1" | jq '.data | {id, title, channelTitle}'
echo -e "\n"

# Test 4: Get video by ID (valid - another video)
echo "4. Get video by ID (ID: 5):"
curl -s -X GET "$BASE_URL/api/v1/public/youtube/5" | jq '.data | {id, title, channelTitle}'
echo -e "\n"

# Test 5: Get video by ID (invalid - not found)
echo "5. Get video by ID (ID: 99999 - should fail):"
curl -s -X GET "$BASE_URL/api/v1/public/youtube/99999" | jq '.'
echo -e "\n"

# Test 6: Get video by ID (invalid - bad format)
echo "6. Get video by ID (invalid format - should fail):"
curl -s -X GET "$BASE_URL/api/v1/public/youtube/abc" | jq '.'
echo -e "\n"

# Test 7: Get random video (call multiple times to verify randomness)
echo "7. Get a random video (first call):"
curl -s -X GET "$BASE_URL/api/v1/public/youtube/random" | jq '.data | {id, title, channelTitle}'
echo -e "\n"

echo "8. Get a random video (second call):"
curl -s -X GET "$BASE_URL/api/v1/public/youtube/random" | jq '.data | {id, title, channelTitle}'
echo -e "\n"

echo "9. Get a random video (third call):"
curl -s -X GET "$BASE_URL/api/v1/public/youtube/random" | jq '.data | {id, title, channelTitle}'
echo -e "\n"

# Test 10: Test pagination edge cases
echo "10. Test pagination - page beyond available data:"
curl -s -X GET "$BASE_URL/api/v1/public/youtube?page=10000&limit=10" | jq '.pagination, .data | length'
echo -e "\n"

# Test 11: Test large limit
echo "11. Test large limit (limit 50):"
curl -s -X GET "$BASE_URL/api/v1/public/youtube?page=1&limit=50" | jq '.pagination, .data | length'
echo -e "\n"

# Test 12: Verify video structure
echo "12. Verify video structure (checking all fields):"
curl -s -X GET "$BASE_URL/api/v1/public/youtube/1" | jq '.data | {id, title, description, channelTitle, publishedAt}'
echo -e "\n"

echo "=== All tests completed ==="
