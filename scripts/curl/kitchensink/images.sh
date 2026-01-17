#!/bin/bash

# Kitchen Sink - Images Test Script
BASE_URL="${BASE_URL:-http://localhost:8080}"

echo "=== Kitchen Sink - Images Tests ==="
echo "Base URL: $BASE_URL"
echo ""

# Test 1: JPEG
echo "1. JPEG Image (Checking Header):"
curl -s -I "$BASE_URL/api/v1/kitchen-sink/images/jpeg" | grep "Content-Type"
echo ""

# Test 2: PNG
echo "2. PNG Image (Checking Header):"
curl -s -I "$BASE_URL/api/v1/kitchen-sink/images/png" | grep "Content-Type"
echo ""

# Test 3: SVG
echo "3. SVG Image (Checking Header):"
curl -s -I "$BASE_URL/api/v1/kitchen-sink/images/svg" | grep "Content-Type"
echo ""

# Test 4: WebP
echo "4. WebP Image (Checking Header):"
curl -s -I "$BASE_URL/api/v1/kitchen-sink/images/webp" | grep "Content-Type"
echo ""

echo "=== All tests completed ==="
