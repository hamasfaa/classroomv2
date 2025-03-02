package main

import (
	"be/config"
	"be/handlers"
	"be/middlewares"
	"be/repositories"
	"be/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(middlewares.SessionMiddleware())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000, http://localhost:5173",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))

	app.Use(logger.New())

	// If you need to auto-migrate your database, uncomment:
	// config.Migration()

	db, _ := config.ConnectDB()
	jwtSecret := "your-secret-key"

	authenticationRepository := repositories.NewAuthenticationRepository(db)
	dosenRepository := repositories.NewDosenRepository(db)
	authenticationHandler := handlers.NewAuthenticationHandler(authenticationRepository, jwtSecret)
	dosenHandler := handlers.NewDosenHandler(dosenRepository)

	routes.SetupRoutes(app, authenticationHandler, dosenHandler, jwtSecret)

	app.Listen(":3000")
}
