package migration

import (
	"github.com/tamnk74/todolist-mysql-go/models"
	"gorm.io/gorm"
)

// Migration for table items, models Item
func MigrateItem(db *gorm.DB) {
	// db.Migrator().DropTable(&models.Item{})
	db.AutoMigrate(&models.Item{})
	// db.Migrator().AddColumn(&models.Item{}, "role")
	// db.Migrator().DropColumn(&models.Item{}, "role")
}
