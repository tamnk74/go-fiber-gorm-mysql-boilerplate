package items

import (
	"github.com/tamnk74/todolist-mysql-go/repository"
)

func NewItemService(itemRepo repository.IItemRepository) *ItemService {
	return &ItemService{
		itemRepo: itemRepo,
	}
}

func NewItemController(itemService IItemService) *ItemController {
	return &ItemController{
		itemService: itemService,
	}
}
