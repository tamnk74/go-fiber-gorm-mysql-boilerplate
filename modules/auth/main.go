package auth

import (
	"github.com/tamnk74/todolist-mysql-go/repository"
)

func NewAuthService(authRepo repository.IAuthRepository) *AuthService {
	return &AuthService{
		authRepo: authRepo,
	}
}

func NewAutController(authService IAuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}
