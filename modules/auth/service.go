package auth

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/tamnk74/todolist-mysql-go/constants"
	"github.com/tamnk74/todolist-mysql-go/models"
	"github.com/tamnk74/todolist-mysql-go/repository"
	Jwt "github.com/tamnk74/todolist-mysql-go/utils/jwt"
	"github.com/tamnk74/todolist-mysql-go/utils/redis"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Login(email string, password string) (token string, err error)
	CreateUser(user models.User) (usr models.User, err error)
	GetUser(user *models.User) error
}

type AuthService struct {
	authRepo repository.IAuthRepository
}

func (a *AuthService) Login(email string, password string) (token string, err error) {
	user, err := a.authRepo.FindUserByEmail(email)
	if err != nil || user.Status == constants.STATUS.INACTIVE {
		return "", errors.New("Invalid email")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println(user.Password, password, err)
		return "", errors.New("Invalid password")
	}

	token = Jwt.GenerateAccessToken(user)
	redis.SaveToken(strconv.Itoa(int(user.ID)), token)
	return token, nil
}

func (a *AuthService) CreateUser(user models.User) (usr models.User, err error) {
	return a.authRepo.Create(user)
}

func (a *AuthService) GetUser(user *models.User) error {
	return a.authRepo.FindByPk(user)
}
