#!/bin/bash

# Auth - Register Test Script
BASE_URL="${BASE_URL:-http://localhost:8080}"
TIMESTAMP=$(date +%s)
USERNAME="user_$TIMESTAMP"
EMAIL="user_$TIMESTAMP@example.com"
PASSWORD="password123"

echo "=== Auth - Register Tests ==="
echo "Base URL: $BASE_URL"
echo ""

# Test 1: Successful Registration
echo "1. Registering new user ($USERNAME)..."
curl -s -X POST "$BASE_URL/api/v1/auth/register" \
  -H "Content-Type: application/json" \
  -d "{
    \"username\": \"$USERNAME\",
    \"email\": \"$EMAIL\",
    \"password\": \"$PASSWORD\"
  }" | jq .
echo ""

# Test 2: Duplicate User
echo "2. Trying to register duplicate user..."
curl -s -X POST "$BASE_URL/api/v1/auth/register" \
  -H "Content-Type: application/json" \
  -d "{
    \"username\": \"$USERNAME\",
    \"email\": \"$EMAIL\",
    \"password\": \"$PASSWORD\"
  }" | jq .
echo ""

# Test 3: Invalid Input
echo "3. Testing invalid input..."
curl -s -X POST "$BASE_URL/api/v1/auth/register" \
  -H "Content-Type: application/json" \
  -d "{
    \"username\": \"ab\",
    \"email\": \"invalid-email\",
    \"password\": \"123\"
  }" | jq .
echo ""

echo "=== All tests completed ==="
