#!/bin/bash
set -e

cd /home/ultimatum/projects_experiments/free_api/apihub_go

# Add all new files
git add .

# Create commit
git commit -m "feat(public): add Random Users API with tests and curl script

- Add RandomUser domain model with complete data structure
- Implement RandomUserService with pagination support
- Create HTTP handler with Swagger annotations
- Add routes for GET /api/v1/public/randomusers (paginated list)
- Add route for GET /api/v1/public/randomusers/:id (get by ID)
- Add route for GET /api/v1/public/randomusers/random (get random user)
- Include comprehensive unit tests (all passing)
- Add curl test script with 10 test cases
- Copy complete randomuser.json data (500+ users)

Tests: 3 test suites, 10 test cases, all passing
API Endpoints: 3 endpoints with pagination and validation"

echo "Commit created successfully!"
git log --oneline -1
git show --stat
