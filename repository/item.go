package repository

import (
	"gorm.io/gorm"

	"github.com/tamnk74/todolist-mysql-go/database"
	dto "github.com/tamnk74/todolist-mysql-go/dto"
	models "github.com/tamnk74/todolist-mysql-go/models"
)

type ItemRepository interface {
	ListItems(pagi *dto.Pagination) (res []models.Item, err error)
	CreateItem(item models.Item) (models.Item, error)
}

type itemRepository struct {
	Conn *gorm.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewItemRepository() ItemRepository {
	return &itemRepository{database.GetDB()}
}

func (m *itemRepository) ListItems(pagi *dto.Pagination) (res []models.Item, err error) {
	var items []models.Item
	var count int64
	m.Conn.Limit(pagi.Limit).Offset(pagi.Offset).Find(&items)
	m.Conn.Model(&models.Item{}).Count(&count)
	pagi.Total = count

	return items, nil
}

func (m *itemRepository) CreateItem(item models.Item) (res models.Item, err error) {
	m.Conn.Create(&item)
	m.Conn.Last(&item)

	return item, nil
}
