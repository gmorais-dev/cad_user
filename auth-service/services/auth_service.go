package services

import (
	"auth-service/dtos"
	"auth-service/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

// Registro de usuário
func (s *AuthService) Register(dto *dtos.RegisterDTO) (*models.User, error) {
	// Verifica se e-mail já está em uso
	var existing models.User
	if err := s.db.Where("email = ?", dto.Email).First(&existing).Error; err == nil {
		return nil, errors.New("email já está em uso")
	}

	// Criptografa a senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("erro ao criptografar senha")
	}

	// Cria o usuário
	user := models.User{
		Username: dto.Username,
		Email:    dto.Email,
		Password: string(hashedPassword),
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Login de usuário
func (s *AuthService) Login(email, password string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("usuário não encontrado")
	}

	// Compara senha
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("senha incorreta")
	}

	return &user, nil
}
