# ✈️ Flight Ticket Booking System (Backend Only)

This is a backend-only project for booking flight tickets. It supports user and admin roles with JWT authentication.

---

## 🔧 Tech Stack

- Go (Gin Framework)
- PostgreSQL
- JWT Authentication
- GORM (ORM)

---

## 📌 Features

- User Signup & Login (with JWT)
- Role-based Access (Admin/User)
- Add/View Flights (Admin/User)
- Book Flight Tickets
- Manage Passengers & Transactions

---

## 🚀 How to Run

1. **Install dependencies**
   ```bash
   go mod tidy

# .env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=777mis
DB_NAME=flight_db
JWT_SECRET=your_jwt_secret

# Run the project
go run main.go

Author
Built by [Mudassar]