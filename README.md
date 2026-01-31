# Go-API-Integration


Go-API-Integration
Overview

This project demonstrates API integration in Golang using a clean architecture approach.

We integrate the public API
👉 https://dummyjson.com/users

and fetch user data into our Go application.


Features:
1) External REST API integration
2) Clean Repository Pattern
3) HTTP client usage in Go
4) JSON response binding into structs
5) Fetching only required fields from API response
6) Error handling and response validation

Go-API-Integration/

│
├── handler/        # HTTP handlers (Gin)

├── service/        # Business logic

├── repository/     # API integration logic

├── model/          # Request/Response models

├── main.go         # Application entry point


How It Works:
1) The handler receives the request.
2) The service layer calls the repository.
3) The repository fetches data from the DummyJSON API.
4) The response is decoded into Go structs.
5) Only required user fields are returned to the client.


How to Run:

go mod tidy
go run main.go
