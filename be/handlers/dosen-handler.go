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

func (h *DosenHandler) GetDetailClass(c *fiber.Ctx) error {
	classID := c.Params("id")

	if classID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID kelas harus diisi"})
	}

	users, err := h.dosenRepository.GetAllUserInClass(classID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	class, err := h.dosenRepository.GetDetailClass(classID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Success",
		"data":    users,
		"class":   class,
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

		filePath := fmt.Sprintf("uploads/tasks/%s/%s", task.TDJudul, fileData.TFNama)

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

func (h *DosenHandler) DeleteTask(c *fiber.Ctx) error {
	taskID := c.Params("id")

	if taskID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID tugas harus diisi"})
	}

	taskDetail, err := h.dosenRepository.GetTaskByID(taskID)
	if err != nil {
		fmt.Printf("Warning: gagal mendapatkan detail task: %v\n", err)
	}

	if err := h.dosenRepository.DeleteTaskWithFiles(taskID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if taskDetail.TDJudul != "" {
		dirPath := fmt.Sprintf("uploads/tasks/%s", taskDetail.TDJudul)
		if err := os.RemoveAll(dirPath); err != nil {
			fmt.Printf("Warning: gagal menghapus direktori %s: %v\n", dirPath, err)
		}
	}

	return c.JSON(fiber.Map{
		"message": "Tugas berhasil dihapus",
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

func (h *DosenHandler) UpdateStatusTask(c *fiber.Ctx) error {
	var status entities.Status

	if err := c.BodyParser(&status); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	taskID := c.Params("id")

	if taskID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID tugas harus diisi"})
	}

	if err := h.dosenRepository.UpdateStatusTask(taskID, status.Status); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Status tugas berhasil diubah",
	})
}

func (h *DosenHandler) CreateMeeting(c *fiber.Ctx) error {
	meeting := new(entities.AbsenDosen)

	if err := c.BodyParser(meeting); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if meeting.ADDeskripsi == "" || meeting.ADPertemuan == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Semua field harus diisi"})
	}

	classID := c.Params("id")
	if classID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "KelasID tidak ditemukan"})
	}
	meeting.KelasKID = classID

	newMUID, err := uuid.NewRandom()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate UUID"})
	}
	meeting.ADID = newMUID.String()

	const charset = "0123456789"

	rand.Seed(time.Now().UnixNano())
	kode := make([]byte, 6)
	for i := 0; i < 6; i++ {
		kode[i] = charset[rand.Intn(len(charset))]
	}
	meeting.ADKode = string(kode)

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
	meeting.UserUID = userUID

	if err := h.dosenRepository.CreateMeeting(meeting); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Pertemuan berhasil dibuat",
	})
}
