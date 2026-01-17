#!/bin/bash
set -e

cd /home/ultimatum/projects_experiments/free_api/apihub_go

# Initialize Go module
go mod init github.com/ultimatum/apihub_go

# Create directory structure
mkdir -p cmd/server
mkdir -p internal/domain internal/repository internal/service internal/handler internal/middleware internal/websocket
mkdir -p pkg/config pkg/logger pkg/database pkg/response pkg/errors pkg/utils
mkdir -p api/swagger
mkdir -p tests/unit tests/integration tests/testdata
mkdir -p scripts/curl scripts/seed
mkdir -p data public/images public/temp .github/workflows

# Create subdirectories
mkdir -p internal/domain/user internal/domain/todo internal/domain/ecommerce internal/domain/social internal/domain/chat
mkdir -p internal/repository/mongodb
mkdir -p internal/service/auth internal/service/todo internal/service/ecommerce internal/service/social internal/service/chat
mkdir -p internal/handler/public internal/handler/auth internal/handler/todo internal/handler/ecommerce internal/handler/social internal/handler/chat internal/handler/kitchen_sink
mkdir -p scripts/curl/public scripts/curl/auth scripts/curl/todo scripts/curl/ecommerce scripts/curl/social scripts/curl/chat scripts/curl/kitchen_sink scripts/curl/seeding

# Create .gitkeep files
touch public/images/.gitkeep public/temp/.gitkeep

echo "Project structure created successfully!"
tree -L 3 -d
