#!/bin/bash
set -e

cd /home/ultimatum/projects_experiments/free_api/apihub_go

echo "Copying complete randomuser data..."
cp /home/ultimatum/projects_experiments/free_api/apihub/src/json/randomuser.json data/randomuser.json

echo "Copying complete randomjoke data..."
cp /home/ultimatum/projects_experiments/free_api/apihub/src/json/randomjoke.json data/randomjoke.json

echo "Running go mod tidy..."
go mod tidy

echo "Running tests..."
go test ./tests/unit/... -v

echo "Building application..."
mkdir -p bin
go build -o bin/apihub_go ./cmd/server

echo "Build successful!"
ls -lh bin/apihub_go
