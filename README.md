# Food Delivery Order API
Project structure:
All backend source code is inside the /backend directory.
📄 Docs conatains

- Detailed design explanation is available in `backend/docs/design_doc.md`
- AI prompts used for development are documented in `backend/docs/prompts_log.md`

A production-ready RESTful Food Delivery Order API built with **Go (Gin)**, **GORM**, **SQLite**, and **JWT authentication**.

---

## 🚀 Live API Endpoints

The backend is structured to handle everything a food delivery platform needs: users, restaurants, menu items, and orders.

### Public Routes
- `GET /ping` - Health check
- `GET /restaurants` - List all restaurants
- `GET /restaurants/:id` - Get a specific restaurant and its menu
- `POST /register` - Register a new user
- `POST /login` - Login and receive a JWT token

### Protected Routes (Requires JWT Token)
*Send token in header: `Authorization: Bearer <token>`*
- `POST /orders` - Place a new order
- `GET /orders/history` - View your order history
- `GET /orders/:id` - Get details of a specific order

### Admin Routes (Requires Admin JWT Token)
- `POST /admin/restaurants` - Add a new restaurant
- `POST /admin/restaurants/:id/menu` - Add menu items to a restaurant
- `PATCH /admin/orders/:id/status` - Update an order's status (`pending` -> `confirmed` -> `preparing` -> `delivered`)

---

## 🛠️ Tech Stack
- **Language**: Go 1.24
- **Framework**: Gin
- **Database**: SQLite (GORM Auto-migration)
- **Authentication**: JWT & bcrypt

---

## 💻 Quickstart (Run Locally)

1. **Clone the repository:**
   ```bash
   git clone https://github.com/Shivani19-code/food-delivery-order-api.git
   cd food-delivery-order-api/backend
   ```

2. **Run the server:**
   ```bash
   go run ./cmd/api/
   ```
   *The server will start on `http://localhost:8080`. The database (`food_delivery.db`) will automatically be created and seeded with sample restaurants!*

3. **Test with curl (Login):**
   ```bash
   curl -X POST http://localhost:8080/register \
     -H "Content-Type: application/json" \
     -d '{"name":"Shivani","email":"shivani@example.com","password":"password123"}'
   ```
