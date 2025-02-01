package main

import (
	"be/config"
	"be/handlers"
	"be/middlewares"
	"be/repositories"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func main() {
	app := fiber.New()

	store := session.New(session.Config{
		CookieHTTPOnly: true,
		CookieSecure:   false,
		Expiration:     time.Hour * 24,
	})

	app.Use(func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Terjadi kesalahan",
			})
		}

		c.Locals("session", sess)
		return c.Next()
	})

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

	api.Get("/protected", func(c *fiber.Ctx) error {
		sess := c.Locals("session").(*session.Session)

		uid := sess.Get("uid")
		if uid == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Belum login atau session tidak valid"})
		}
		return c.JSON(fiber.Map{"message": "Anda sudah login", "uid": uid})
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
