package main

import (
	"be/config"
	"be/handlers"
	"be/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, http://localhost:5173",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))
	app.Use(logger.New())

	// config.Migration()

	db, _ := config.ConnectDB()

	authenticationRepository := repositories.NewAuthenticationRepository(db)

	jwtSecret := "your-secret-key"

	authenticationHandler := handlers.NewAuthenticationHandler(authenticationRepository, jwtSecret)

	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api.Post("/register", authenticationHandler.RegisterUser)
	api.Post("/login", authenticationHandler.LoginUser)

	app.Listen(":3000")

}
