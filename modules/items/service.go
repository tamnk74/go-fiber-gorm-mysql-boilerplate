package items

import (
	"github.com/tamnk74/todolist-mysql-go/dto"
	"github.com/tamnk74/todolist-mysql-go/models"
	"github.com/tamnk74/todolist-mysql-go/repository"
)

type IItemService interface {
	ListItems(pagi *dto.Pagination) ([]models.Item, error)
	CreateItem(item models.Item) (models.Item, error)
}

type ItemService struct {
	itemRepo repository.IItemRepository
}

func (a *ItemService) ListItems(pagi *dto.Pagination) (res []models.Item, err error) {
	return a.itemRepo.ListItems(pagi)
}

func (a *ItemService) CreateItem(item models.Item) (res models.Item, err error) {
	return a.itemRepo.CreateItem(item)
}
