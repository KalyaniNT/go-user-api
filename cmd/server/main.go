package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"go-user-api/config"
	"go-user-api/internal/handler"
	"go-user-api/internal/repository"
	"go-user-api/internal/routes"
	"go-user-api/internal/service"
)

func main() {
	app := fiber.New()

	// DB connection
	db := config.ConnectDB()

	// Dependency injection
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Routes
	routes.RegisterUserRoutes(app, userHandler)

	log.Fatal(app.Listen(":3000"))
}
