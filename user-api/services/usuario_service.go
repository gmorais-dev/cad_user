package services

import (
	"errors"
	"user-api/config"
	"user-api/dtos"
	"user-api/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// hashPassword gera o hash da senha com bcrypt
func hashPassword(senha string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkEmailDuplicado(email string, ignoreID uint64) (bool, error) {
	var usuario models.Usuario
	query := config.DB.Where("email = ?", email)
	if ignoreID != 0 {
		query = query.Where("id != ?", ignoreID)
	}
	err := query.First(&usuario).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return err == nil, err
}

func AtualizarUsuario(id uint64, dto dtos.UsuarioRequest) (*models.Usuario, error) {
	var usuario models.Usuario

	if err := config.DB.First(&usuario, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("usuário não encontrado")
		}
		return nil, err
	}

	duplicado, err := checkEmailDuplicado(dto.Email, id)
	if err != nil {
		return nil, err
	}
	if duplicado {
		return nil, errors.New("email já cadastrado")
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

	duplicado, err := checkEmailDuplicado(dto.Email, 0)
	if err != nil {
		return nil, err
	}
	if duplicado {
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
	var usuario models.Usuario
	if err := config.DB.First(&usuario, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("usuário não encontrado")
		}
		return err
	}

	if err := config.DB.Delete(&usuario).Error; err != nil {
		return err
	}
	return nil
}
