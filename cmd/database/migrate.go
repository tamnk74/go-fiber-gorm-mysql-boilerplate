package main

import (
	"os"

	"github.com/tamnk74/todolist-mysql-go/database"
	"github.com/tamnk74/todolist-mysql-go/database/migration"
	"github.com/tamnk74/todolist-mysql-go/database/seeder"
)

func main() {
	err := database.Connect()
	if err != nil {
		panic("Failed to connect database")
	}
	db := database.GetDB()

	migration.MigrateUser(db)
	migration.MigrateItem(db)

	if len(os.Args) <= 2 {
		return
	}
	option := os.Args[1]
	if option == "seed" {
		switch os.Args[2] {
		case "users":
			seeder.SeedUser(db)
		case "items":
			seeder.SeedItem(db)
		}
	}
}
