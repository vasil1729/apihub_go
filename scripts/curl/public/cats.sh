#!/bin/bash
BASE_URL="${BASE_URL:-http://localhost:8080}"
echo "=== Cats API Tests ==="
curl -s "$BASE_URL/api/v1/public/cats" | jq '.data | length'
curl -s "$BASE_URL/api/v1/public/cats/1" | jq '.data.name'
curl -s "$BASE_URL/api/v1/public/cats/random" | jq '.data.name'
echo "=== All tests completed ==="
