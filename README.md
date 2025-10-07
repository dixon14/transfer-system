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

3. **The API will be available at:** `http://localhost:8080`

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

4. **Install dependencies:**
   ```bash
   make deps
   # OR
   go mod download
   ```

5. **Run the application:**
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

## Deployment

### Deploy to Remote Server

1. **Build and push Docker image:**
   ```bash
   # Build the image
   docker build -t your-registry/transfer-system:latest .

   # Push to registry
   docker push your-registry/transfer-system:latest
   ```

2. **On the remote server, create docker-compose.yml:**
   ```yaml
   version: '3.8'

   services:
     postgres:
       image: postgres:15-alpine
       environment:
         POSTGRES_USER: postgres
         POSTGRES_PASSWORD: <secure-password>
         POSTGRES_DB: transfer_system
       volumes:
         - postgres_data:/var/lib/postgresql/data
         - ./migrations/init.sql:/docker-entrypoint-initdb.d/init.sql
       restart: always

     app:
       image: your-registry/transfer-system:latest
       environment:
         DB_HOST: postgres
         DB_PORT: 5432
         DB_USER: postgres
         DB_PASSWORD: <secure-password>
         DB_NAME: transfer_system
         DB_SSLMODE: disable
         PORT: 8080
       ports:
         - "8080:8080"
       depends_on:
         - postgres
       restart: always

   volumes:
     postgres_data:
   ```

3. **Start the services:**
   ```bash
   docker-compose up -d
   ```

### Deploy as Standalone Binary

1. **Build for target platform:**
   ```bash
   # For Linux
   GOOS=linux GOARCH=amd64 go build -o transfer-system main.go

   # For Windows
   GOOS=windows GOARCH=amd64 go build -o transfer-system.exe main.go
   ```

2. **Transfer binary to server and set environment variables**

3. **Run with systemd (Linux):**
   Create `/etc/systemd/system/transfer-system.service`:
   ```ini
   [Unit]
   Description=Transfer System API
   After=network.target postgresql.service

   [Service]
   Type=simple
   User=appuser
   WorkingDirectory=/opt/transfer-system
   Environment="DB_HOST=localhost"
   Environment="DB_PORT=5432"
   Environment="DB_USER=postgres"
   Environment="DB_PASSWORD=<secure-password>"
   Environment="DB_NAME=transfer_system"
   Environment="PORT=8080"
   ExecStart=/opt/transfer-system/transfer-system
   Restart=on-failure

   [Install]
   WantedBy=multi-user.target
   ```

   Enable and start:
   ```bash
   sudo systemctl enable transfer-system
   sudo systemctl start transfer-system
   ```

**Health Check:**
```bash
curl -X GET http://localhost:8080/health
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
