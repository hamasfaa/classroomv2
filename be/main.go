package main

import (
	"be/config"
	"be/handlers"
	"be/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	// config.Migration()

	db, _ := config.ConnectDB()

	authenticationRepository := repositories.NewAuthenticationRepository(db)

	jwtSecret := "your-secret-key"

	authenticationHandler := handlers.NewAuthenticationHandler(authenticationRepository, jwtSecret)

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/register", authenticationHandler.RegisterUser)
	app.Post("/login", authenticationHandler.LoginUser)

	app.Listen(":3000")

}
