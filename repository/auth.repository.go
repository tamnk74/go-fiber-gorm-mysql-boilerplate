package repository

import (
	"errors"

	"gorm.io/gorm"

	models "github.com/tamnk74/todolist-mysql-go/models"
)

type IAuthRepository interface {
	FindByPk(*models.User) error
	FindUserByEmail(email string) (res models.User, err error)
	Create(user models.User) (res models.User, err error)
}

type AuthRepository struct {
	Conn *gorm.DB
}

func (m *AuthRepository) FindUserByEmail(email string) (res models.User, err error) {
	var user models.User
	result := m.Conn.Where("email = ?", email).First(&user)
	if result.RowsAffected == 0 {
		return models.User{}, errors.New("Incorrect email")
	}
	return user, nil
}

func (m *AuthRepository) FindByPk(user *models.User) error {
	result := m.Conn.First(&user)
	if result.RowsAffected == 0 {
		return errors.New("User not found")
	}
	return nil
}

func (m *AuthRepository) Create(user models.User) (res models.User, err error) {
	result := m.Conn.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}
