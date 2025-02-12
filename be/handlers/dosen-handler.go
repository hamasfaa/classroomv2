package handlers

import (
	"be/entities"
	"be/repositories"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type DosenHandler struct {
	dosenRepository repositories.DosenRepository
}

func NewDosenHandler(dosenRepository repositories.DosenRepository) *DosenHandler {
	return &DosenHandler{dosenRepository: dosenRepository}
}

func (h *DosenHandler) CreateClass(c *fiber.Ctx) error {
	class := new(entities.Kelas)
	if err := c.BodyParser(class); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if class.KNamaKelas == "" || class.KMataKuliah == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Semua field harus diisi"})
	}

	newCUID, err := uuid.NewRandom()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate UUID",
		})
	}
	class.KID = newCUID.String()

	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	rand.Seed(time.Now().UnixNano())
	kode := make([]byte, 6)
	for i := 0; i < 6; i++ {
		kode[i] = charset[rand.Intn(len(charset))]
	}
	class.KKodeKelas = string(kode)

	class.KTanggalDibuat = time.Now()

	if err := h.dosenRepository.CreateClass(class); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	userToken := c.Locals("user")
	if userToken == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	token := userToken.(*jwt.Token)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
	}

	userUID, ok := claims["uid"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
	}

	if err := h.dosenRepository.AddUserToClass(userUID, class.KID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Kelas berhasil dibuat",
	})
}

func (h *DosenHandler) GetAllClass(c *fiber.Ctx) error {
	userToken := c.Locals("user")
	if userToken == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	token := userToken.(*jwt.Token)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
	}

	userUID, ok := claims["uid"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
	}

	classes, err := h.dosenRepository.GetAllClass(userUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Success",
		"data":    classes,
	})
}

func (h *DosenHandler) DeleteClass(c *fiber.Ctx) error {
	classID := c.Params("id")

	if classID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID kelas harus diisi"})
	}

	if err := h.dosenRepository.DeleteClass(classID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Kelas berhasil dihapus",
	})
}
