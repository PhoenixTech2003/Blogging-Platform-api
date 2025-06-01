# Blogging Platform API 📝

[![Go Version](https://img.shields.io/badge/Go-1.24.3-00ADD8?style=for-the-badge&logo=go)](https://go.dev/)
[![Docker](https://img.shields.io/badge/Docker-Containerized-2496ED?style=for-the-badge&logo=docker)](https://www.docker.com/)
[![CI/CD](https://img.shields.io/badge/CI%2FCD-GitHub_Actions-2088FF?style=for-the-badge&logo=github-actions)](https://github.com/PhoenixTech2003/Blogging-Platform-api/actions)
[![Swagger](https://img.shields.io/badge/Swagger-API_Docs-85EA2D?style=for-the-badge&logo=swagger)](http://localhost:8081/swagger/index.html)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Database-336791?style=for-the-badge&logo=postgresql)](https://www.postgresql.org/)
[![RESTful](https://img.shields.io/badge/RESTful-API-009688?style=for-the-badge&logo=fastapi)]()

A modern RESTful API for a blogging platform built with Go. This API allows you to create and retrieve blog articles with a clean, well-documented interface.

## 🚀 Features

- Create new blog articles with title and content
- Retrieve all articles or filter by search query
- RESTful API design
- Swagger documentation
- PostgreSQL database for data persistence

## 📁 Project Structure

```
├── docs/                  # Swagger documentation files
├── internal/              # Internal application code
│   ├── database/          # Database connection and generated queries
│   │   ├── articles.sql.go  # Generated SQL code for articles
│   │   ├── db.go           # Database connection setup
│   │   └── models.go       # Database models
│   ├── handlers/          # HTTP request handlers
│   │   └── articles.go    # Article-related handlers
│   ├── router/            # HTTP routing
│   │   └── articles.go    # Article routes
│   └── utils/             # Utility functions
├── scripts/              # Helper scripts
├── sql/                  # SQL files
│   ├── queries/          # SQL query definitions
│   │   └── articles.sql  # Article queries
│   └── schema/           # Database schema
│       └── 001_article.sql # Article table schema
├── .env                  # Environment variables (gitignored)
├── go.mod                # Go module dependencies
├── go.sum                # Go module checksums
└── main.go              # Application entry point
```

## 🛠️ Technologies Used

- **Go** - Programming language
- **PostgreSQL** - Database
- **Docker** - Containerization
- **Swagger** - API documentation
- **sqlc** - SQL query generation
- **goose** - Database migrations

## 🔧 Prerequisites

- Go 1.24.3 or higher
- PostgreSQL database
- Git

## ⚙️ Environment Setup

Create a `.env` file in the root directory with the following variables:

```env
DATABASE_URL=postgresql://username:password@localhost:5432/blogdb?sslmode=disable
BASE_URL=http://localhost:8081
SWAGGER_HOST=localhost:8081
```

## 📥 Installation

1. Clone the repository
   ```bash
   git clone https://github.com/PhoenixTech2003/Blogging-Platform-api.git
   cd Blogging-Platform-api
   ```

2. Install dependencies
   ```bash
   go mod download
   ```

3. Set up the database
   ```bash
   # Install goose for migrations
   go install github.com/pressly/goose/v3/cmd/goose@latest
   
   # Run migrations
   goose -dir sql/schema postgres "your-database-connection-string" up
   ```

4. Run the application
   ```bash
   go run main.go
   ```

5. Access the Swagger documentation at http://localhost:8081/swagger/index.html



## 🧪 Development

### Generate SQL code

This project uses [sqlc](https://sqlc.dev/) to generate type-safe Go code from SQL:

```bash
# Install sqlc
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# Generate code
sqlc generate
```

### Update Swagger Documentation

```bash
# Install swag
go install github.com/swaggo/swag/cmd/swag@latest

# Generate swagger docs
swag init
```
## 👥 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
