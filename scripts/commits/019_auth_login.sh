#!/bin/bash
set -e

# Change to the root directory of the project
cd "$(dirname "$0")/../.."

# Tidy modules
echo "Tidying modules..."
go mod tidy

# Generate Swagger docs (silently)
echo "Generating Swagger docs..."
swag init -g cmd/server/main.go -o docs > /dev/null 2>&1 || true

# Run Unit Tests
echo "Running Unit Tests..."
go test -v ./tests/unit/auth

# Run Integration Tests
echo "Running Integration Tests..."
go test -v ./tests/integration/auth_test.go ./tests/integration/setup_test.go

# Commit
echo "Committing..."
git add .
git commit -m "🔐 feat(auth): add User Login API" -m "- Implemented User Login endpoint (POST /api/v1/auth/login)" -m "- Added JWT token generation using 'golang-jwt/jwt/v5'" -m "- Added LoginRequest model" -m "- Added Unit tests for Login" -m "- Added Integration tests (mocked) for Login"
