package routes

import (
	"be/handlers"
	"be/middlewares"

	"github.com/gofiber/fiber/v2"
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

	api.Get("/protected", authenticationHandler.Whoami)

	apiDosen.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Dosen!")
	})
	apiDosen.Post("/addClass", dosenHandler.CreateClass)
	apiDosen.Get("/class", dosenHandler.GetAllClass)
	apiDosen.Delete("/deleteClass/:id", dosenHandler.DeleteClass)
	apiDosen.Post("/addTask/:id", dosenHandler.CreateTask)
	apiDosen.Get("/manageTask/:id", dosenHandler.GetAllTask)
	apiDosen.Put("/updateTaskStatus/:id", dosenHandler.UpdateStatusTask)
	apiDosen.Get("/detailClass/:id", dosenHandler.GetDetailClass)

	apiMahasiswa.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Mahasiswa!")
	})
}
