package repository

import "github.com/tamnk74/todolist-mysql-go/database"

func NewAuthRepository() IAuthRepository {
	return &AuthRepository{database.GetDB()}
}

func NewItemRepository() IItemRepository {
	return &ItemRepository{database.GetDB()}
}
