#!/bin/bash
set -e
cd /home/ultimatum/projects_experiments/free_api/apihub_go
echo "Copying data files..."
cp /home/ultimatum/projects_experiments/free_api/apihub/src/json/{randomuser,randomjoke,quotes,nse-stocks,meals,dogs,cats}.json data/
go mod tidy
go test ./tests/unit/... -v
go test ./tests/integration/... -v
mkdir -p bin
go build -o bin/apihub_go ./cmd/server
ls -lh bin/apihub_go
