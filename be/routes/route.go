package routes

import (
	"be/handlers"
	"be/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func SetupRoutes(app *fiber.App, authenticationHandler *handlers.AuthenticationHandler, dosenHandler *handlers.DosenHandler, jwtSecret string) {
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
	apiDosen.Get("/class", dosenHandler.GetAllClass)
	apiDosen.Delete("/deleteClass/:id", dosenHandler.DeleteClass)
	apiDosen.Post("/addTask/:id", dosenHandler.CreateTask)

	apiMahasiswa.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Mahasiswa!")
	})
}
