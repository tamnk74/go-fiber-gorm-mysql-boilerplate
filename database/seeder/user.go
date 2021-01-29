package seeder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tamnk74/todolist-mysql-go/models"
	"gorm.io/gorm"
)

// Migration for table users, models User
func SeedUser(db *gorm.DB) {
	jsonFile, err := os.Open("database/data/users.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Opened users.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users []models.User
	json.Unmarshal([]byte(byteValue), &users)
	db.Create(&users)
}
