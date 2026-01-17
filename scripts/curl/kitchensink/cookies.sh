#!/bin/bash

# Kitchen Sink - Cookies Test Script
BASE_URL="${BASE_URL:-http://localhost:8080}"
COOKIE_JAR="cookies.txt"

echo "=== Kitchen Sink - Cookies Tests ==="
echo "Base URL: $BASE_URL"
echo ""

# Test 1: Set Cookie
echo "1. Set Cookie (foo=bar):"
curl -s -c $COOKIE_JAR -i "$BASE_URL/api/v1/kitchen-sink/cookies/set?key=foo&value=bar" | grep "Set-Cookie"
echo "Check cookie jar:"
cat $COOKIE_JAR
echo ""

# Test 2: Get Cookies
echo "2. Get Cookies:"
curl -s -b $COOKIE_JAR "$BASE_URL/api/v1/kitchen-sink/cookies/get" | jq '.data'
echo ""

# Test 3: Delete Cookie
echo "3. Delete Cookie (foo):"
curl -s -c $COOKIE_JAR -i "$BASE_URL/api/v1/kitchen-sink/cookies/delete?key=foo" | grep "Set-Cookie"
echo ""

# Cleanup
rm $COOKIE_JAR

echo "=== All tests completed ==="
