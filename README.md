# Banking System Backend

## Project Overview

The **Banking System Backend** is a Go-based project that manages the core functionalities of a banking application. It handles secure user management, account operations, transaction management, employee/role management, and reporting.

### Technology Stack:

- **Backend**: Go (Gin framework)
- **Database**: PostgreSQL (GORM for ORM)
- **Authentication**: JWT (JSON Web Token)
- **Configuration**: YAML for database settings

## Project Structure

```bash
banking-system-backend/
├── application/          # Business logic layer
├── database/             # tables schema
├── docs/                 # Api docs
├── domain/               # Domain models/entities
├── inbound/              # API handlers (controllers)
├── outbound/             # Database and external services
├── util/                 # Utilities and helper functions
├── config.yaml           # Configuration file
├── main.go               # Main entry point
├── go.mod                # Go module file
└── go.sum                # Go dependencies lock file
```

## Api Documentation

for api documentation please click here [API Guide](docs/banking-system-backend.postman_collection.json)

## Setup and Installation

### 1. Clone the repository:

```bash
git clone https://github.com/Niteshkumarjain-ui/banking-system-backend.git
cd banking-system-backend
```

### 2. Install Go dependencies:

```bash
go mode tidy
```

### 3. Setup PostgreSQL:

- Create PostgresSql database
- Update the config.yaml file with your PostgreSQL credentials
- Update your jwt key in config.yaml file
- For table schema click [here](database/)

### 4. Setup Jaeger

```bash
docker run -d --name jaeger -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p 5775:5775/udp -p 6831:6831/udp -p 6832:6832/udp -p 5778:5778 -p 16686:16686 -p 14268:14268 -p 9411:9411 jaegertracing/all-in-one:1.6
```

### 5. Run the application:

```bash
go run main.go
```

### 6. Access the server:

The server will run on http://localhost:8000.
