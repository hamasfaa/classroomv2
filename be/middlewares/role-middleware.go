package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func MahasiswaOnly(secretKey string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, ok := c.Locals("session").(*session.Session)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Session tidak ditemukan",
			})
		}

		role, ok := sess.Get("role").(string)
		if !ok || role != "mahasiswa" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Hanya untuk role mahasiswa",
			})
		}

		return c.Next()
	}
}

func DosenOnly(secretKey string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, ok := c.Locals("session").(*session.Session)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Session tidak ditemukan",
			})
		}

		role, ok := sess.Get("role").(string)
		if !ok || role != "dosen" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Hanya untuk role dosen",
			})
		}

		return c.Next()
	}
}
