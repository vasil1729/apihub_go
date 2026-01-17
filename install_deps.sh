#!/bin/bash
set -e

cd /home/ultimatum/projects_experiments/free_api/apihub_go

echo "Installing Go dependencies..."

# Core dependencies
go get github.com/gin-gonic/gin
go get github.com/gin-contrib/cors
go get github.com/rs/zerolog
go get github.com/joho/godotenv
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options
go get go.mongodb.org/mongo-driver/mongo/readpref
go get go.mongodb.org/mongo-driver/bson
go get go.mongodb.org/mongo-driver/bson/primitive

# JWT
go get github.com/golang-jwt/jwt/v5

# Password hashing
go get golang.org/x/crypto/bcrypt

# Validation
go get github.com/go-playground/validator/v10

# UUID
go get github.com/google/uuid

# Testing
go get github.com/stretchr/testify

# Tidy up
go mod tidy

echo "Dependencies installed successfully!"
