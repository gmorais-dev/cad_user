package services

import (
	"errors"
	"user-api/config"
	"user-api/dtos"
	"user-api/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func hashPassword(senha string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	return string(bytes), err
}

func AtualizarUsuario(id uint64, dto dtos.UsuarioRequest) (*models.Usuario, error) {
	var usuario models.Usuario

	if err := config.DB.First(&usuario, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("usuário não encontrado")
		}
		return nil, err
	}

	if usuario.Email != dto.Email {
		var outro models.Usuario
		if err := config.DB.Where("email = ?", dto.Email).First(&outro).Error; err == nil {
			return nil, errors.New("email já cadastrado")
		}
	}

	usuario.Nome = dto.Nome
	usuario.Email = dto.Email

	if dto.Senha != "" {
		hashed, err := hashPassword(dto.Senha)
		if err != nil {
			return nil, errors.New("erro ao criptografar senha")
		}
		usuario.Senha = hashed
	}

	if err := config.DB.Save(&usuario).Error; err != nil {
		return nil, err
	}

	return &usuario, nil
}

func CriarUsuario(dto dtos.UsuarioRequest) (*models.Usuario, error) {
	var usuarioExistente models.Usuario
	if err := config.DB.Where("email = ?", dto.Email).First(&usuarioExistente).Error; err == nil {
		return nil, errors.New("email já cadastrado")
	}

	hashed, err := hashPassword(dto.Senha)
	if err != nil {
		return nil, errors.New("erro ao criptografar senha")
	}

	usuario := models.Usuario{
		Nome:  dto.Nome,
		Email: dto.Email,
		Senha: hashed,
	}

	if err := config.DB.Create(&usuario).Error; err != nil {
		return nil, err
	}

	return &usuario, nil
}

func ListarUsuarios(limit, offset int) ([]models.Usuario, error) {
	var usuarios []models.Usuario
	if err := config.DB.Limit(limit).Offset(offset).Find(&usuarios).Error; err != nil {
		return nil, err
	}
	return usuarios, nil
}
func BuscarUsuarioPorID(id uint64) (*models.Usuario, error) {
	var usuario models.Usuario
	if err := config.DB.First(&usuario, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("usuário não encontrado")
		}
		return nil, err
	}
	return &usuario, nil
}

func DeletarUsuario(id uint64) error {
	if err := config.DB.Delete(&models.Usuario{}, id).Error; err != nil {
		return err
	}
	return nil
}
