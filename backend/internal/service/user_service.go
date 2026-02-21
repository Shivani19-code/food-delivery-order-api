package service

import (
	"errors"
	"food-delivery-api/internal/model"
	"food-delivery-api/pkg/config"
	"food-delivery-api/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct{}

func (s *UserService) Register(user *model.User) error {
	// Check if user already exists
	var existingUser model.User
	if err := config.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return errors.New("user with this email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return config.DB.Create(user).Error
}

func (s *UserService) Login(email, password string) (string, *model.User, error) {
	var user model.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil, errors.New("invalid email or password")
		}
		return "", nil, err
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", nil, err
	}

	// Remove password before returning
	user.Password = ""
	return token, &user, nil
}
