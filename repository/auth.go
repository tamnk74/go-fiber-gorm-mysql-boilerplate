package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/tamnk74/todolist-mysql-go/database"
	models "github.com/tamnk74/todolist-mysql-go/models"
)

type AuthRepository interface {
	FindByPk(*models.User) error
	FindUserByEmail(email string) (res models.User, err error)
	Create(user models.User) (res models.User, err error)
}

type authRepository struct {
	Conn *gorm.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewAuthRepository() AuthRepository {
	return &authRepository{database.GetDB()}
}

func (m *authRepository) FindUserByEmail(email string) (res models.User, err error) {
	var user models.User
	result := m.Conn.Where("email = ?", email).First(&user)
	if result.RowsAffected == 0 {
		return models.User{}, errors.New("Incorrect email")
	}
	return user, nil
}

func (m *authRepository) FindByPk(user *models.User) error {
	result := m.Conn.First(&user)
	if result.RowsAffected == 0 {
		return errors.New("User not found")
	}
	return nil
}

func (m *authRepository) Create(user models.User) (res models.User, err error) {
	result := m.Conn.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}
