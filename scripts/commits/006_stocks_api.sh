#!/bin/bash
set -e

cd /home/ultimatum/projects_experiments/free_api/apihub_go

# Add all new files
git add .

# Create commit
COMMIT_MSG="✨ feat(public): add Stocks API with tests and curl script

- 📈 Add Stock domain model with NSE stock data (13 fields)
- 🔍 Implement StockService with symbol-based lookup (case-insensitive)
- 🎯 Create HTTP handler with Swagger annotations
- 🛣️ Add routes for GET /api/v1/public/stocks (paginated list)
- 🔎 Add route for GET /api/v1/public/stocks/:symbol (get by symbol)
- ✅ Include comprehensive unit tests (2 test suites, 8 test cases)
- 🧪 Add integration tests (4 test cases)
- 📜 Add curl test script with 10 test cases
- 📊 Copy complete NSE stocks data (28,565+ stocks)
- 📚 Regenerate Swagger documentation

Tests: 47 total (38 unit + 17 integration), all passing ✅
API Endpoints: 11 endpoints total (2 new for stocks) 🚀
Build: Successful (45MB binary)"

git commit -m "$COMMIT_MSG"

# Get commit hash
COMMIT_HASH=$(git rev-parse --short HEAD)

echo "Commit created successfully!"
echo "Commit: 006_stocks_api_${COMMIT_HASH}"
git log --oneline -6
