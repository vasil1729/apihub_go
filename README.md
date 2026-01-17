# FreeAPI.app - Go Edition

[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-ISC-blue.svg)](LICENSE.md)

## Overview

**apihub_go** is a comprehensive API hub built with Go, providing a wide range of APIs for learning and development. This is a Go implementation of the popular [FreeAPI project](https://github.com/hiteshchoudhary/apihub), featuring clean architecture, test-driven development, and production-ready code.

### Key Features

- 🏗️ **Clean Architecture** - Clear separation of concerns with domain, repository, service, and handler layers
- ✅ **Test-Driven Development** - Comprehensive unit and integration tests
- 📝 **API Documentation** - Auto-generated Swagger/OpenAPI documentation
- 🔒 **Secure** - JWT authentication, bcrypt password hashing, rate limiting
- 🚀 **Production-Ready** - Docker support, graceful shutdown, structured logging
- 🎯 **Multiple API Categories** - Public APIs, Authentication, E-commerce, Social Media, Chat, and more

## API Categories

1. **Public APIs** - Random users, jokes, quotes, stocks, meals, pets, books, YouTube data
2. **Kitchen Sink** - HTTP methods, status codes, request/response inspection, cookies, redirects, images
3. **Authentication** - Registration, login, JWT tokens, password management, email verification
4. **Todo List** - Full CRUD operations for todo management
5. **E-commerce** - Products, cart, orders, payments, addresses, coupons
6. **Social Media** - Posts, comments, likes, bookmarks, follow system
7. **Chat App** - One-on-one and group chats with WebSocket support
8. **Database Seeding** - Utilities to populate the database with test data

## Quick Start

### Prerequisites

- Go 1.24 or higher
- MongoDB 7.0 or higher
- Docker & Docker Compose (optional)

### Installation

#### Using Docker (Recommended)

```bash
# Clone the repository
git clone https://github.com/ultimatum/apihub_go.git
cd apihub_go

# Copy environment file
cp .env.sample .env

# Edit .env and set your secrets
nano .env

# Start with Docker Compose
docker-compose up --build
```

#### Running Locally

```bash
# Clone the repository
git clone https://github.com/ultimatum/apihub_go.git
cd apihub_go

# Install dependencies
go mod download

# Copy environment file
cp .env.sample .env

# Edit .env and set your configuration
nano .env

# Run the application
go run cmd/server/main.go

# Or use Make
make run
```

### Using Makefile

```bash
make help              # Show all available commands
make run               # Run the application
make build             # Build the binary
make test              # Run all tests
make test-unit         # Run unit tests only
make test-integration  # Run integration tests only
make coverage          # Generate coverage report
make lint              # Run linter
make swagger           # Generate Swagger docs
make docker-up         # Start Docker containers
make docker-down       # Stop Docker containers
```

## Configuration

All configuration is done through environment variables. Copy `.env.sample` to `.env` and configure:

```env
# Server
PORT=8080
NODE_ENV=development

# Database
MONGODB_URI=mongodb://localhost:27017/apihub_go
DB_NAME=apihub_go

# JWT Secrets (CHANGE THESE!)
ACCESS_TOKEN_SECRET=your-super-secret-access-token
REFRESH_TOKEN_SECRET=your-super-secret-refresh-token

# CORS
CORS_ORIGIN=http://localhost:3000,http://localhost:5173

# Rate Limiting
RATE_LIMIT_WINDOW_MS=900000
RATE_LIMIT_MAX_REQUESTS=5000
```

## API Documentation

Once the server is running, access the Swagger documentation at:

```
http://localhost:8080/swagger/index.html
```

## Testing

### Run All Tests

```bash
make test
```

### Run Unit Tests

```bash
make test-unit
```

### Run Integration Tests

```bash
make test-integration
```

### Generate Coverage Report

```bash
make coverage
```

This will generate `coverage.html` which you can open in your browser.

## Project Structure

```
apihub_go/
├── cmd/
│   └── server/          # Application entry point
├── internal/
│   ├── domain/          # Business entities
│   ├── repository/      # Data access layer
│   ├── service/         # Business logic
│   ├── handler/         # HTTP handlers
│   ├── middleware/      # HTTP middleware
│   └── websocket/       # WebSocket handlers
├── pkg/
│   ├── config/          # Configuration
│   ├── logger/          # Logging
│   ├── database/        # Database connection
│   ├── response/        # API responses
│   ├── errors/          # Custom errors
│   └── utils/           # Utilities
├── tests/
│   ├── unit/            # Unit tests
│   ├── integration/     # Integration tests
│   └── testdata/        # Test fixtures
├── scripts/
│   ├── curl/            # Curl test scripts
│   └── seed/            # Database seeding
├── data/                # Static JSON data
└── public/              # Uploaded files
```

## Architecture

This project follows **Clean Architecture** principles:

1. **Domain Layer** - Pure business logic, no external dependencies
2. **Repository Layer** - Database operations
3. **Service Layer** - Business logic orchestration
4. **Handler Layer** - HTTP request/response handling
5. **Middleware Layer** - Cross-cutting concerns

## Development Workflow

Each API endpoint follows Test-Driven Development:

1. Write tests first (unit + integration)
2. Implement the feature
3. Create curl test script
4. Commit (one commit per endpoint)

## Curl Test Scripts

Each API endpoint has a corresponding curl script in `scripts/curl/`:

```bash
# Example: Test user registration
bash scripts/curl/auth/register.sh

# Run all curl tests
bash scripts/curl/run_all.sh
```

## Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

### Contribution Guidelines

1. Fork the repository
2. Create a feature branch: `git checkout -b feat/your-feature`
3. Write tests for your changes
4. Ensure all tests pass: `make test`
5. Run linter: `make lint`
6. Commit your changes: `git commit -am 'feat: add your feature'`
7. Push to the branch: `git push origin feat/your-feature`
8. Submit a pull request

## License

This project is licensed under the ISC License - see the [LICENSE.md](LICENSE.md) file for details.

## Acknowledgments

- Inspired by the original [FreeAPI project](https://github.com/hiteshchoudhary/apihub) by Hitesh Choudhary
- Built with ❤️ using Go and modern best practices

## Support

- 📧 Email: support@apihub.com
- 🐛 Issues: [GitHub Issues](https://github.com/ultimatum/apihub_go/issues)
- 💬 Discussions: [GitHub Discussions](https://github.com/ultimatum/apihub_go/discussions)

## Roadmap

- [x] Project setup and core infrastructure
- [ ] Public APIs implementation
- [ ] Kitchen Sink APIs
- [ ] Authentication system
- [ ] Todo List APIs
- [ ] E-commerce APIs
- [ ] Social Media APIs
- [ ] Chat App with WebSocket
- [ ] Complete test coverage
- [ ] CI/CD pipeline
- [ ] Deployment guides

---

**Made with Go 🚀**
