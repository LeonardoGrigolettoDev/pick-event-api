package services

import (
	"errors"

	"github.com/LeonardoGrigolettoDev/pick-point.git/database"
	"github.com/LeonardoGrigolettoDev/pick-point.git/models"
	"github.com/LeonardoGrigolettoDev/pick-point.git/utils"
)

// Registro de usuário
func RegisterUser(user *models.User) (string, error) {
	//TODO before inserting user in DB, it have to be created Entity, this is about fk dependencies.
	// Hash da senha antes de salvar
	if err := user.HashPassword(); err != nil {
		return "", err
	}

	// Salvar usuário no banco
	if err := database.DB.Create(user).Error; err != nil {
		return "", err
	}

	// Gerar token JWT
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Login de usuário
func LoginUser(email, password string) (string, error) {
	var user models.User

	// Buscar usuário pelo e-mail
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("usuário não encontrado")
	}

	// Verificar senha
	if !user.CheckPassword(password) {
		return "", errors.New("senha incorreta")
	}

	// Gerar token JWT
	return utils.GenerateToken(user.ID)
}
