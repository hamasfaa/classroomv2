package repositories

import (
	"be/entities"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthenticationRepository interface {
	CreateUser(user *entities.User) error
	AuthenticationUser(email, password string) (*entities.User, error)
}

type authenticationRepositoryGorm struct {
	db *gorm.DB
}

func NewAuthenticationRepository(db *gorm.DB) *authenticationRepositoryGorm {
	return &authenticationRepositoryGorm{db: db}
}

func (r *authenticationRepositoryGorm) CreateUser(user *entities.User) error {
	var existingUser entities.User
	err := r.db.Where("u_email = ?", user.UEmail).First(&existingUser).Error

	if err == nil {
		return errors.New("email already exists")
	} else if err != gorm.ErrRecordNotFound {
		return err
	}

	return r.db.Create(user).Error
}

func (r *authenticationRepositoryGorm) AuthenticationUser(email, password string) (*entities.User, error) {
	var user entities.User

	if err := r.db.Where("u_email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.UPassword), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	return &user, nil
}
