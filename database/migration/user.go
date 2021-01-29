package migration

import (
	"github.com/tamnk74/todolist-mysql-go/models"
	"gorm.io/gorm"
)

// Migration for table users, models User
func MigrateUser(db *gorm.DB) {
	// db.Migrator().DropTable(&models.User{})
	db.AutoMigrate(&models.User{})
	// db.Migrator().AddColumn(&models.User{}, "role")
	// db.Migrator().DropColumn(&models.User{}, "role")
}
