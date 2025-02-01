package handlers

import (
	"be/entities"
	"be/repositories"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
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

	sess := c.Locals("session").(*session.Session)
	userUID := sess.Get("uid").(string)
	if err := h.dosenRepository.AddUserToClass(userUID, class.KID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Kelas berhasil dibuat",
	})
}
