#!/bin/bash
set -e

# Generate Swagger docs
echo "Generating Swagger docs..."
~/go/bin/swag init -g cmd/server/main.go -o ./docs

# Run tests
echo "Running Unit Tests..."
go test ./tests/unit/kitchensink/... -v

echo "Running Integration Tests..."
go test ./tests/integration/kitchensink_images_test.go ./tests/integration/setup_test.go -v

# Commit (NO PUSH)
echo "Committing..."
git add .
git commit -m "🖼️ feat(kitchen-sink): add Images API

- Implemented GET /kitchen-sink/images/jpeg: Generates dynamic random JPEG
- Implemented GET /kitchen-sink/images/png: Generates dynamic random PNG
- Implemented GET /kitchen-sink/images/svg: Returns a static SVG image
- Implemented GET /kitchen-sink/images/webp: Returns a static WebP image (1x1 pixel)
- Added unit and integration tests for all image formats
- Created usage example curl script"
git log -1
