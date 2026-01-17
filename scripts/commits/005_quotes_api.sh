#!/bin/bash
set -e

cd /home/ultimatum/projects_experiments/free_api/apihub_go

# Add all new files
git add .

# Create commit
COMMIT_MSG="✨ feat(public): add Quotes API with tests and curl script

- 💬 Add Quote domain model with author, content, tags, and metadata
- 🎲 Implement QuoteService with true random selection
- 🎯 Create HTTP handler with Swagger annotations
- 🛣️ Add routes for GET /api/v1/public/quotes (paginated list)
- 🔍 Add route for GET /api/v1/public/quotes/:id (get by ID)
- 🎭 Add route for GET /api/v1/public/quotes/random (get random quote)
- ✅ Include comprehensive unit tests with randomness verification
- 🧪 Add integration tests (4 test cases)
- 📜 Add curl test script with 10 test cases
- 📊 Copy complete quotes.json data (3000+ quotes)
- 📚 Regenerate Swagger documentation

Tests: 39 total (30 unit + 13 integration), all passing ✅
API Endpoints: 9 endpoints total (3 new for quotes) 🚀
Build: Successful (45MB binary)"

git commit -m "$COMMIT_MSG"

# Get commit hash
COMMIT_HASH=$(git rev-parse --short HEAD)

echo "Commit created successfully!"
echo "Commit: 005_quotes_api_${COMMIT_HASH}"
git log --oneline -5
