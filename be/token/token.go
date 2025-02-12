package token

import (
	"be/entities"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *entities.User, jwtSecret string) (string, error) {
	claims := jwt.MapClaims{
		"uid":   user.UID,
		"email": user.UEmail,
		"role":  user.URole,
		"exp":   time.Now().Add(time.Minute * 5).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtSecret))
}

func GenerateRefreshToken(user *entities.User, jwtSecret string) (string, error) {
	claims := jwt.MapClaims{
		"uid": user.UID,
		"exp": time.Now().Add(time.Hour * 24 * 14).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}
