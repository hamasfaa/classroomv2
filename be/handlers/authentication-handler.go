package handlers

import (
	"be/entities"
	"be/repositories"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationHandler struct {
	authenticationRepository repositories.AuthenticationRepository
	JWTSecret                string
}

func NewAuthenticationHandler(authenticationRepository repositories.AuthenticationRepository, jwtsecret string) *AuthenticationHandler {
	return &AuthenticationHandler{authenticationRepository: authenticationRepository, JWTSecret: jwtsecret}
}

func (h *AuthenticationHandler) RegisterUser(c *fiber.Ctx) error {
	user := new(entities.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if user.UNama == "" || user.UEmail == "" || user.UPassword == "" || user.URole == "" || user.UTanggalLahir.IsZero() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Semua field harus diisi"})
	}

	newUUID, err := uuid.NewRandom()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate UUID",
		})
	}
	user.UID = newUUID.String()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.UPassword), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}
	user.UPassword = string(hashedPassword)

	if err := h.authenticationRepository.CreateUser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user.UPassword = ""
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Registrasi berhasil",
		"user":    user,
	})
}

func (h *AuthenticationHandler) LoginUser(c *fiber.Ctx) error {
	type loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	input := new(loginRequest)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	user, err := h.authenticationRepository.AuthenticationUser(input.Email, input.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	token, err := h.GenerateJWT(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	sess := c.Locals("session").(*session.Session)
	sess.Set("uid", user.UID)
	sess.Set("role", user.URole)
	if err := sess.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"error": "Gagal menyimpan session"},
		)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token":   token,
		"message": "Login berhasil, session disimpan",
	})
}

func (h *AuthenticationHandler) LogoutUser(c *fiber.Ctx) error {
	sess := c.Locals("session").(*session.Session)
	err := sess.Destroy()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal menghapus session",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Logout berhasil, session telah dihapus",
	})
}

func (h *AuthenticationHandler) GenerateJWT(user *entities.User) (string, error) {
	claims := jwt.MapClaims{
		"uid":   user.UID,
		"email": user.UEmail,
		"role":  user.URole,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(h.JWTSecret))
}
