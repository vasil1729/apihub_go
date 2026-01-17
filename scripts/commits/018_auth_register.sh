#!/bin/bash
set -e

# Generate Swagger docs
echo "Generating Swagger docs..."
~/go/bin/swag init -g cmd/server/main.go -o ./docs

# Run tests
echo "Running Unit Tests..."
go test ./tests/unit/auth/... -v

echo "Running Integration Tests..."
go test ./tests/integration/auth_register_test.go ./tests/integration/setup_test.go -v

# Commit (NO PUSH)
echo "Committing..."
git add .
git commit -m "🔐 feat(auth): add User Registration API

- Implemented User Registration endpoint (POST /api/v1/auth/register)
- Added 'bcrypt' for password hashing
- Created User model with JSON/BSON tags
- Used MongoDB for user persistence
- Added 'mtest' unit tests for AuthService
- Added integration tests verifying success, duplicate, and invalid input scenarios"
git log -1
