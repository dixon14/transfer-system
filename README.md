# Transfer System - Financial Transaction API

A robust financial transaction system built with Golang, Gin framework, and PostgreSQL, featuring proper MVC architecture and ACID-compliant transactions.

## Features

- **Account Management**: Create and query accounts with initial balances
- **Fund Transfers**: Secure transfers between accounts with data integrity
- **ACID Compliance**: Database transactions ensure consistency
- **Input Validation**: Comprehensive request validation and error handling
- **Clean Architecture**: MVC structure with clear separation of concerns

## Prerequisites

- Go 1.24 or higher
- PostgreSQL 15+ (or use Docker)
- Docker & Docker Compose (for containerized deployment)
- Make (for using Makefile commands)

## Assumptions
-  Currency is not considered here, but can be easily implemented in both the transactions and accounts table.
- Authentication can be implemented with the convenience of Gin's middleware
- Index should be properly implemented in transactions table if there is a need to frequently query transaction records

## Running Locally

### Option 1: Using Docker Compose (Recommended)

1. **Clone and navigate to the project:**
   ```bash
   cd transfer-system
   ```

2. **Start the application:**
   ```bash
   make docker-run
   # OR
   docker-compose up --build
   ```

3. **The API will be available at:** `http://localhost:8080` and you can curl the `http://localhost:8080/health` endpoint to check if the application is up and running.

4. **Stop the application:**
   ```bash
   make docker-stop
   # OR
   docker-compose down
   ```

### Option 2: Running Manually

1. **Start PostgreSQL:**
   ```bash
   docker run -d \
     --name postgres \
     -e POSTGRES_USER=postgres \
     -e POSTGRES_PASSWORD=postgres \
     -e POSTGRES_DB=transfer_system \
     -p 5432:5432 \
     postgres:15-alpine
   ```

2. **Initialize the database:**
   ```bash
   psql -h localhost -U postgres -d transfer_system -f migrations/init.sql
   ```

3. **Set environment variables:**
   ```bash
   export DB_HOST=localhost
   export DB_PORT=5432
   export DB_USER=postgres
   export DB_PASSWORD=postgres
   export DB_NAME=transfer_system
   export DB_SSLMODE=disable
   export PORT=8080
   ```

4. **Run the application:**
   ```bash
   make run
   # OR
   go run main.go
   ```

## Building the Application

### Build Binary
```bash
make build
```
The binary will be created at `bin/transfer-system`

### Clean Build Artifacts
```bash
make clean
```
## Configuration

Environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `DB_HOST` | PostgreSQL host | `localhost` |
| `DB_PORT` | PostgreSQL port | `5432` |
| `DB_USER` | Database user | `postgres` |
| `DB_PASSWORD` | Database password | `postgres` |
| `DB_NAME` | Database name | `transfer_system` |
| `DB_SSLMODE` | SSL mode | `disable` |
| `PORT` | Application port | `8080` |
| `GIN_MODE` | Gin mode (debug/release) | `release` |

## Data Integrity & Consistency

The system ensures data integrity through:

1. **Database Transactions**: All fund transfers use PostgreSQL transactions with row-level locking (`SELECT FOR UPDATE`)
2. **Constraint Checks**: Balance cannot be negative, amounts must be positive
3. **Validation Layer**: Input validation at controller and service layers
4. **Error Handling**: Comprehensive error handling with proper rollback mechanisms
5. **Atomic Operations**: Transfers are atomic - either both accounts are updated or neither

## Error Handling

The API returns appropriate HTTP status codes:

- `200 OK`: Successful operation
- `400 Bad Request`: Invalid input or business logic error
- `500 Internal Server Error`: Server error

All errors include a JSON response with `error` and `message` fields.
