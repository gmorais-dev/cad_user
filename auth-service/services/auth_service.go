package services

import (
	"errors"

	"auth-service/config"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func NewAuthService() *AuthService {
	return &AuthService{
		DB: config.DB,
	}
}

func (s *AuthService) Register(user *models.User) error {
	// Verifica se email já existe
	var existingUser models.User
	if err := s.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return errors.New("email already registered")
	}

	// Hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	// Salvar no banco
	return s.DB.Create(user).Error
}

func (s *AuthService) Login(email, password string) (*models.User, error) {
	var user models.User
	if err := s.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("invalid credentials")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}
