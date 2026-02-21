# Design Document: Food Delivery Order API

## 1. Introduction
The Food Delivery Order API is designed to provide a backend for a modern food delivery application. It handles user authentication, restaurant management, menu browsing, and order processing.

## 2. Architectural Design
The project follows **Clean Architecture** patterns to separate concerns and ensure testability:

- **Handler Layer**: Responsible for handling HTTP requests, parsing JSON, and returning appropriate HTTP status codes.
- **Service Layer**: Contains the core business logic (e.g., calculating total order price, verifying stock).
- **Model Layer**: Defines the data structures and GORM tags for database interaction.
- **Repository Layer**: (Abstracted by GORM) Handles data persistence.

## 3. Data Models
### User
- Role-based access control (Admin/Customer).
- Encrypted passwords using `bcrypt`.

### Restaurant & MenuItem
- One-to-Many relationship between Restaurant and MenuItems.
- Restaurants have metadata like rating and address.

### Order & OrderItem
- Orders track the state (Pending -> Preparing -> Delivered).
- OrderItems store a snapshot of the price at the time of purchase to ensure auditability even if menu prices change later.

## 4. Key Design Decisions
### 4.1. JWT Authentication
We chose JSON Web Tokens (JWT) for authentication because they are stateless, making the API scalable and easy to integrate with modern frontend frameworks or mobile apps.

### 4.2. Database Transactions
Order placement uses **GORM Transactions**. This ensures that either the entire order (including all items) is saved, or none of it is, preventing "partial orders" in case of system failure.

### 4.3. SQLite for Portability
For the purpose of evaluation and rapid prototyping, SQLite was chosen as it requires zero setup for the evaluator while providing full SQL capabilities.

## 5. Security Considerations
- **Password Hashing**: Never store plain text passwords.
- **Protected Routes**: Middleware verifies JWT before allowing sensitive actions.
- **Role Verification**: Admin routes are restricted using a specialized middleware.

## 6. Future Improvements
- Implement Redis for caching frequent restaurant searches.
- Add real-time order tracking using WebSockets.
- Integrate a real payment gateway (e.g., Stripe/Razorpay).
