# Design and Implementation Reasoning

## Overall Architecture
The application follows a clean layered architecture to ensure separation of concerns and maintainability.

- **Handler layer**: Handles HTTP requests and responses using Fiber.
- **Service layer**: Contains business logic such as age calculation.
- **Repository layer**: Manages database interactions.
- **SQLC**: Generates type-safe database access code.

## Database Design
The `users` table stores the user's name and date of birth (DOB).  
Age is not stored in the database to avoid inconsistency; instead, it is calculated dynamically when required.

## Age Calculation
Age is calculated using Go’s `time` package by comparing the current date with the DOB.  
This ensures accuracy across leap years and date boundaries.

## SQLC Usage
SQLC was used to generate strongly typed queries from SQL files.  
This reduces runtime SQL errors and improves code safety and readability.

## Input Validation
The `go-playground/validator` library is used to validate incoming request data, ensuring required fields are present and correctly formatted.

## Logging
Uber Zap is used for structured logging, especially during database connection and server startup, to aid debugging and observability.

## Design Decisions
- Dynamic age calculation instead of storing age
- Clean separation between layers
- SQLC for type-safe database access

This approach results in a scalable, readable, and maintainable backend service.
