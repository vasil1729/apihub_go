#!/bin/bash
set -e
cd /home/ultimatum/projects_experiments/free_api/apihub_go

# Add books init and routes
cat >> internal/handler/public/routes.go << 'EOF'

	// Initialize book service and handler
	bookService, err := publicService.NewBookService(dataPath)
	if err != nil {
		return err
	}
	bookHandler := NewBookHandler(bookService)
	
	// Books routes
	books := router.Group("/books")
	{
		books.GET("", bookHandler.GetAll)
		books.GET("/random", bookHandler.GetRandom)
		books.GET("/:id", bookHandler.GetByID)
	}
EOF

# Copy data
cp /home/ultimatum/projects_experiments/free_api/apihub/src/json/books.json data/

# Build and test
go mod tidy
go test ./... -v -short
go build -o bin/apihub_go ./cmd/server

# Swagger and commit
~/go/bin/swag init -g cmd/server/main.go -o ./docs
git add .
git commit -m "📚 feat(public): add Books API (8/10 complete)"
git log --oneline -1
