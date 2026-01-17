#!/bin/bash
echo "Testing User Login..."

# Register a user first (ignore error if exists)
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "loginuser",
    "email": "login@example.com",
    "password": "password123"
  }' > /dev/null 2>&1

# Login
echo -e "\n\nLogging in..."
curl -v -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "login@example.com",
    "password": "password123"
  '
