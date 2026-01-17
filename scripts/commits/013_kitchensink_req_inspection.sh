#!/bin/bash
set -e

# Generate Swagger docs
echo "Generating Swagger docs..."
~/go/bin/swag init -g cmd/server/main.go -o ./docs

# Run tests
echo "Running Unit Tests..."
go test ./tests/unit/kitchensink/... -v

echo "Running Integration Tests..."
go test ./tests/integration/kitchensink_request_inspection_test.go ./tests/integration/setup_test.go -v

# Commit
echo "Committing..."
git add .
git commit -m "🕵️ feat(kitchen-sink): add Request Inspection API"
git log --oneline -1
