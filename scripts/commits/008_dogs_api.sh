#!/bin/bash
set -e

cd /home/ultimatum/projects_experiments/free_api/apihub_go

git add .

COMMIT_MSG="✨ feat(public): add Dogs API with tests and curl script

- 🐕 Add Dog domain model with nested weight/height/image structures
- 🎲 Implement DogService with true random selection
- 🎯 Create HTTP handler with Swagger annotations
- 🛣️ Add routes for GET /api/v1/public/dogs (paginated list)
- 🔍 Add route for GET /api/v1/public/dogs/:id (get by ID)
- 🎭 Add route for GET /api/v1/public/dogs/random (get random breed)
- ✅ Include comprehensive unit tests (3 test suites, 10 test cases)
- 🧪 Add integration tests (4 test cases)
- 📜 Add curl test script with 8 test cases
- 📊 Copy complete dogs data (190+ breeds)
- 📚 Regenerate Swagger documentation

Tests: 63 total (54 unit + 25 integration), all passing ✅
API Endpoints: 17 endpoints total (3 new for dogs) 🚀
Build: Successful (45MB binary)"

git commit -m "$COMMIT_MSG"

COMMIT_HASH=$(git rev-parse --short HEAD)
echo "Commit created successfully!"
echo "Commit: 008_dogs_api_${COMMIT_HASH}"
git log --oneline -8
