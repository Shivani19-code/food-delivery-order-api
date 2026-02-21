# Food Delivery Order API

A robust, production-grade RESTful API for a food delivery service built using Go, Gin, and GORM. This project was developed as part of the Infosys Go Project evaluation.

## 🚀 Key Features

- **User Authentication**: Secure JWT-based registration and login with password hashing (bcrypt).
- **Restaurant Management**: CRUD operations for restaurants and their menus.
- **Order Flow**: Multi-item order placement with atomic transactions and status tracking.
- **Auto-Seeding**: The database is automatically populated with sample data on the first run for immediate testing.
- **Clean Architecture**: Organized into Handlers, Services, and Models for better maintainability.
- **SQLite Database**: Lightweight and portable database setup.

## 🛠️ Technology Stack

- **Language**: Go (Golang) 1.21+
- **Framework**: [Gin Gonic](https://github.com/gin-gonic/gin)
- **ORM**: [GORM](https://gorm.io/)
- **Database**: SQLite3
- **Auth**: JWT (JSON Web Tokens)

## 🏃 How to Run

1. **Prerequisites**: Ensure Go is installed on your system.
2. **Clone/Navigate**:
   ```bash
   cd backend
   ```
3. **Install Dependencies**:
   ```bash
   go mod tidy
   ```
4. **Run the Application**:
   ```bash
   go run cmd/api/main.go
   ```
5. **Base URL**: `http://localhost:8080`

## 📡 API Endpoints

### Public Endpoints
- `GET /ping`: Health check.
- `POST /register`: Register a new user (`name`, `email`, `password`).
- `POST /login`: Login to get JWT token (`email`, `password`).
- `GET /restaurants`: List all available restaurants.
- `GET /restaurants/:id`: Get menu for a specific restaurant.

### Protected Endpoints (Requires Bearer Token)
- `POST /orders`: Place a new order with items.
- `GET /orders/history`: View all your past orders.
- `GET /orders/:id`: Track specific order status.

### Admin Endpoints (Requires Admin Role)
- `POST /admin/restaurants`: Add a new restaurant.
- `POST /admin/restaurants/:id/menu`: Add items to a menu.
- `PATCH /admin/orders/:id/status`: Update order status (e.g., delivered).

## 📂 Project Structure

```text
backend/
├── cmd/api/            # Entry point (main.go)
├── internal/
│   ├── handler/        # HTTP Handlers
│   ├── middleware/     # Auth & Logging
│   ├── model/          # Domain Models (GORM)
│   └── service/        # Business Logic
├── pkg/
│   ├── config/         # DB Connection & Seeding
│   └── utils/          # JWT Utilities
└── food_delivery.db    # Generated SQLite DB
```

## 📝 Documentation
Detailed design decisions and implementation details can be found in the [Design Document](./docs/design_doc.md).
All prompts used [Prompts Log](./docs/prompts_log.md).
