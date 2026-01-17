#!/bin/bash
set -e

# Change to the root directory of the project
cd "$(dirname "$0")/../.."

# Generate Swagger docs (silently)
echo "Generating Swagger docs..."
swag init -g cmd/server/main.go -o docs > /dev/null 2>&1 || true

# Run Unit Tests
# No new unit tests for logout as logic is trivial handler only
echo "Running Unit Tests..."
go test -v ./tests/unit/auth

# Run Integration Tests
echo "Running Integration Tests..."
go test -v ./tests/integration/auth_test.go ./tests/integration/setup_test.go

# Commit
echo "Committing..."
git add .
git commit -m "🔐 feat(auth): add User Logout API" -m "- Implemented User Logout endpoint (POST /api/v1/auth/logout)" -m "- Added Integration test for Logout"
