#!/bin/bash
set -e

# Change to the root directory of the project
cd "$(dirname "$0")/../.."

# Generate Swagger docs (silently)
echo "Generating Swagger docs..."
swag init -g cmd/server/main.go -o docs > /dev/null 2>&1 || true

# Run Unit Tests
echo "Running Unit Tests..."
go test -v ./tests/unit/auth

# Run Integration Tests
echo "Running Integration Tests..."
go test -v ./tests/integration/auth_test.go ./tests/integration/auth_refresh_test.go ./tests/integration/setup_test.go

# Commit
echo "Committing..."
git add .
git commit -m "🔐 feat(auth): add Refresh Token API" -m "- Implemented POST /api/v1/auth/refresh" -m "- Updated Login to return Access & Refresh Tokens" -m "- Added Unit & Integration Tests for Refresh Token"
