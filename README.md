# Golang REST API - Feature-Based Architecture

A production-ready, scalable Go REST API boilerplate built using **Gin Framework**, **GORM ORM**, and **JWT Authentication**, structured with a **Feature-Based (Modular / Clean Architecture)** design pattern.

---

## 🌟 Key Features

- **Feature-Based Architecture**: Modules are self-contained (`internal/modules/<feature>`), keeping domain logic isolated and highly maintainable.
- **Clean Layering**: Clear separation of concerns (**Model**, **DTO**, **Repository**, **Service**, **Controller**, **Route**).
- **JWT Authentication & Middleware**: Secure endpoint authorization with stateless JWT tokens.
- **Role-Based Authorization**: Extensible role system (`jwt.Role`) with middleware support.
- **ORM & Auto Migration**: GORM integration with MySQL and automated database migrations.
- **Dependency Injection**: Explicit dependency wiring in `main.go`.

---

## 📁 Project Structure

```text
golang-restapi-big-structure/
├── cmd/
│   └── api/
│       └── main.go               # Application entry point & dependency wiring
├── internal/
│   ├── config/
│   │   └── database.go           # GORM connection & AutoMigrate setup
│   ├── middleware/
│   │   └── auth.go               # JWT Authentication & Role Authorization Middleware
│   ├── pkg/
│   │   └── jwt/                  # Shared JWT service & Claims definition
│   └── modules/                  # Feature-Based Modules
│       ├── user/                 # User Feature Module
│       │   ├── controller.go     # HTTP Handlers (Gin)
│       │   ├── dto.go            # Request Data Transfer Objects & Validation
│       │   ├── model.go          # Database Model Entity
│       │   ├── repository.go     # Database Queries (GORM)
│       │   ├── route.go          # Module Route Definitions
│       │   └── service.go        # Business Logic
│       └── product/              # Product Feature Module
│           ├── controller.go
│           ├── dto.go
│           ├── model.go
│           ├── repository.go
│           ├── route.go
│           └── service.go
├── go.mod
├── go.sum
├── .gitignore
└── README.md
```

---

## 🏗️ Architecture Layers Breakdown

Each module follows a 5-layer architecture:

| Layer | Responsibility |
| :--- | :--- |
| **Model** | Defines the database schema entity for GORM. |
| **DTO** | Data Transfer Objects used to parse and validate incoming JSON payloads (`binding:"required"`). |
| **Repository** | Handles direct database interactions and GORM queries. |
| **Service** | Contains core business logic (e.g., password hashing, default role assignment, business rules). |
| **Controller** | Handles HTTP requests, extracts parameters/JWT context, and returns standard JSON responses. |
| **Route** | Registers module routes (Public & Protected) to Gin's Router Group. |

---

## 🚀 Getting Started

### Prerequisites

- [Go](https://go.dev/doc/install) (v1.20 or later)
- [MySQL](https://www.mysql.com/) server running locally or remotely.

### Database Setup

1. Create a MySQL database (default name: `golang_restapi_big_structure`):
   ```sql
   CREATE DATABASE golang_restapi_big_structure;
   ```
2. Update the Database DSN in [`internal/config/database.go`](file:///c:/Users/user/Desktop/golang-restapi-big-structure/internal/config/database.go) if your credentials differ:
   ```go
   dsn := "root:password@(127.0.0.1:3306)/golang_restapi_big_structure?charset=utf8mb4&Loc=local&parseTime=true"
   ```

---

## 💻 Running the Application

1. **Install dependencies**:
   ```bash
   go mod download
   ```

2. **Run the server**:
   ```bash
   go run ./cmd/api/main.go
   ```
   The server will start at `http://localhost:8080`.

---

## 📡 API Endpoints

### User Module (`/api/v1/users`)

| Method | Endpoint | Access | Description |
| :--- | :--- | :--- | :--- |
| `POST` | `/api/v1/users/register` | Public | Register a new user (defaults to `user` role) |
| `GET` | `/api/v1/users/` | Protected | Get list of all users |
| `GET` | `/api/v1/users/profile` | Protected | Get current user details from JWT token |

### Product Module (`/api/v1/products`)

| Method | Endpoint | Access | Description |
| :--- | :--- | :--- | :--- |
| `GET` | `/api/v1/products/` | Public | Get all products with preloaded creator info |
| `GET` | `/api/v1/products/:id` | Public | Get product detail by ID |
| `POST` | `/api/v1/products/` | Protected | Create a new product attached to current user |

---

## 🔐 Authentication Example

To access protected endpoints, pass the JWT Token in the request header:

```http
Authorization: Bearer <your_jwt_token_here>
```
