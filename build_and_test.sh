#!/bin/bash
set -e

cd /home/ultimatum/projects_experiments/free_api/apihub_go

echo "Copying data files..."
cp /home/ultimatum/projects_experiments/free_api/apihub/src/json/randomuser.json data/
cp /home/ultimatum/projects_experiments/free_api/apihub/src/json/randomjoke.json data/
cp /home/ultimatum/projects_experiments/free_api/apihub/src/json/quotes.json data/
cp /home/ultimatum/projects_experiments/free_api/apihub/src/json/nse-stocks.json data/
cp /home/ultimatum/projects_experiments/free_api/apihub/src/json/meals.json data/
cp /home/ultimatum/projects_experiments/free_api/apihub/src/json/dogs.json data/

echo "Running go mod tidy..."
go mod tidy

echo "Running tests..."
go test ./tests/unit/... -v
go test ./tests/integration/... -v

echo "Building application..."
mkdir -p bin
go build -o bin/apihub_go ./cmd/server

echo "Build successful!"
ls -lh bin/apihub_go
