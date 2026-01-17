#!/bin/bash
set -e

cd /home/ultimatum/projects_experiments/free_api/apihub_go

# Add all new files
git add .

# Create commit
git commit -m "feat(public): add Random Jokes API with tests and curl script

- Add RandomJoke domain model with categories support
- Implement RandomJokeService with true random selection
- Create HTTP handler with Swagger annotations
- Add routes for GET /api/v1/public/randomjokes (paginated list)
- Add route for GET /api/v1/public/randomjokes/:id (get by ID)
- Add route for GET /api/v1/public/randomjokes/random (get random joke)
- Include comprehensive unit tests with randomness verification
- Add curl test script with 10 test cases
- Copy complete randomjoke.json data (1800+ jokes)

Tests: 6 test suites, 20 test cases, all passing
API Endpoints: 6 endpoints total (3 new for jokes)"

echo "Commit created successfully!"
git log --oneline -2
git show --stat
