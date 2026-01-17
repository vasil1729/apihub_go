#!/bin/bash
set -e

cd /home/ultimatum/projects_experiments/free_api/apihub_go

# Add all new files
git add .

# Create commit
COMMIT_MSG="✨ feat(infrastructure): add Swagger UI, integration tests, and coverage reporting

- 📚 Add Swagger UI with auto-generated documentation
- 🔧 Install swag, gin-swagger, and swagger files dependencies
- 🎯 Enable Swagger endpoint at /swagger/index.html
- 🏠 Add root redirect to Swagger UI
- 📊 Generate Swagger docs (docs/swagger.json, swagger.yaml)
- ✅ Create integration test framework for HTTP handlers
- 🧪 Add integration tests for Random Users API (5 test cases)
- 🧪 Add integration tests for Random Jokes API (4 test cases)
- 📈 Update Makefile with swagger generation target
- 📊 Enhance coverage reporting with atomic mode
- 🔧 Update Makefile clean target for docs folder

Tests: 29 total (20 unit + 9 integration), all passing ✅
Coverage: HTML report generated (coverage.html)
Swagger: Accessible at http://localhost:8080/swagger/index.html 🚀"

git commit -m "$COMMIT_MSG"

# Get commit hash
COMMIT_HASH=$(git rev-parse --short HEAD)

echo "Commit created successfully!"
echo "Commit: 004_infrastructure_swagger_tests_${COMMIT_HASH}"
git log --oneline -4
