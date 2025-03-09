package handlers

import (
	"be/entities"
	"be/repositories"
	"be/token"

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
	input := new(entities.Login)
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

	accessToken, err := token.GenerateToken(user, h.JWTSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	refreshToken, err := token.GenerateRefreshToken(user, h.JWTSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate refresh token",
		})
	}

	sess := c.Locals("session").(*session.Session)
	sess.Set("user_id", user.UID)
	if err := sess.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token":         accessToken,
		"refresh_token": refreshToken,
		"message":       "Login berhasil, session disimpan",
	})
}

func (h *AuthenticationHandler) LogoutUser(c *fiber.Ctx) error {
	sess := c.Locals("session").(*session.Session)
	if err := sess.Destroy(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Logout berhasil",
	})
}

func (h *AuthenticationHandler) RefreshToken(c *fiber.Ctx) error {
	input := new(entities.RefreshToken)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid refresh token payload",
		})
	}

	parsedToken, err := jwt.Parse(input.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.JWTSecret), nil
	})
	if err != nil || !parsedToken.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid refresh token",
		})
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid refresh token claims",
		})
	}

	uid, ok := claims["uid"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid refresh token claims",
		})
	}

	user, err := h.authenticationRepository.GetUserByUID(uid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	newToken, err := token.GenerateToken(user, h.JWTSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate new token",
		})
	}

	return c.JSON(fiber.Map{
		"token":   newToken,
		"message": "Refresh token berhasil",
	})
}

func (h *AuthenticationHandler) Whoami(c *fiber.Ctx) error {
	sess := c.Locals("session").(*session.Session)
	userID := sess.Get("user_id")
	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Belum login atau session tidak valid",
		})
	}

	user, err := h.authenticationRepository.GetUserByUID(userID.(string))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	publicUser := entities.PublicUser{
		UNama:         user.UNama,
		UEmail:        user.UEmail,
		URole:         user.URole,
		UTanggalLahir: user.UTanggalLahir,
		UNoPonsel:     user.UNoPonsel,
		UAlamat:       user.UAlamat,
		UFoto:         user.UFoto,
	}

	return c.JSON(fiber.Map{
		"message": "Anda sudah login",
		"user":    publicUser,
	})
}
