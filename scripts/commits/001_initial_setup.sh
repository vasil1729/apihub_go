#!/bin/bash
set -e

cd /home/ultimatum/projects_experiments/free_api/apihub_go

# Add all files
git add .

# Create initial commit
COMMIT_MSG="🎉 chore: initial project setup with core infrastructure

- 📦 Initialize Go module and Git repository
- 🏗️ Set up Clean Architecture directory structure
- ⚙️ Create configuration management with environment variables
- 📝 Implement structured logging with zerolog
- 🗄️ Set up MongoDB connection with health checks
- ✅ Create standardized API response helpers
- ❌ Implement custom error types
- 🛡️ Add middleware: CORS, logging, rate limiting, error handling
- 🚀 Create main server entry point with graceful shutdown
- 🐳 Add Docker and docker-compose configuration
- 🔧 Create Makefile with common development commands
- 📚 Add comprehensive README and LICENSE
- 📥 Install all required dependencies"

git commit -m "$COMMIT_MSG"

# Get commit hash
COMMIT_HASH=$(git rev-parse --short HEAD)

echo "Commit created successfully!"
echo "Commit: 001_initial_setup_${COMMIT_HASH}"
git log --oneline -1
