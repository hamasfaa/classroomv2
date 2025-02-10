package main

import (
	"be/config"
	"be/handlers"
	"be/middlewares"
	"be/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
	app := fiber.New()

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

	api := app.Group("/api")
	apiDosen := api.Group("/dosen", middlewares.DosenOnly(jwtSecret))
	apiMahasiswa := api.Group("/mahasiswa", middlewares.MahasiswaOnly(jwtSecret))

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api.Post("/register", authenticationHandler.RegisterUser)
	api.Post("/login", authenticationHandler.LoginUser)
	api.Post("/logout", authenticationHandler.LogoutUser)
	api.Post("/refreshToken", authenticationHandler.RefreshToken)

	api.Get("/protected", func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Belum login atau token tidak valid"})
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token tidak valid"})
		}
		claims := token.Claims.(jwt.MapClaims)
		return c.JSON(fiber.Map{"message": "Anda sudah login", "uid": claims["uid"]})
	})

	apiDosen.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Dosen!")
	})
	apiDosen.Post("/addClass", dosenHandler.CreateClass)

	apiMahasiswa.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Mahasiswa!")
	})

	app.Listen(":3000")
}
