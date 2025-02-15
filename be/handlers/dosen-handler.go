package handlers

import (
	"be/entities"
	"be/repositories"
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
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

func (h *DosenHandler) CreateTask(c *fiber.Ctx) error {
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

	var request entities.Tugas
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if request.TDJudul == "" || request.TDDeskripsi == "" || request.TDDeadline.IsZero() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Semua field harus diisi"})
	}

	newTUID, err := uuid.NewRandom()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate UUID",
		})
	}

	task := &entities.TugasDosen{
		TDID:            newTUID.String(),
		TDJudul:         request.TDJudul,
		TDDeskripsi:     request.TDDeskripsi,
		TDTanggalDibuat: time.Now(),
		TDDeadline:      request.TDDeadline,
		KelasKID:        c.Params("id"),
		UserUID:         userUID,
	}

	var taskFiles []entities.TugasFile
	for _, fileData := range request.Files {
		fileID, _ := uuid.NewRandom()

		decoded, err := base64.StdEncoding.DecodeString(fileData.TFContent)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid file content",
			})
		}

		filePath := fmt.Sprintf("uploads/tasks/%s/%s", newTUID.String(), fileData.TFNama)

		os.MkdirAll(filepath.Dir(filePath), 0755)

		if err := os.WriteFile(filePath, decoded, 0644); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to save file",
			})
		}

		taskFile := entities.TugasFile{
			TFID:      fileID.String(),
			TFNama:    fileData.TFNama,
			TFPath:    filePath,
			TugasTDID: newTUID.String(),
		}

		taskFiles = append(taskFiles, taskFile)
	}

	if err := h.dosenRepository.CreateTaskWithFiles(task, taskFiles); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Tugas berhasil dibuat",
	})
}

func (h *DosenHandler) GetAllTask(c *fiber.Ctx) error {
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

	classID := c.Params("id")

	if classID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID kelas harus diisi"})
	}

	tasks, err := h.dosenRepository.GetAllTask(userUID, classID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Success",
		"data":    tasks,
	})
}
