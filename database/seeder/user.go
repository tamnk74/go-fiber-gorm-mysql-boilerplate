package seeder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tamnk74/todolist-mysql-go/models"
	"gorm.io/gorm"
)

type UserJSON struct {
	Password string `json:"password"`
}

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
	var userJSONs []UserJSON
	json.Unmarshal([]byte(byteValue), &users)
	json.Unmarshal([]byte(byteValue), &userJSONs)

	for index, user := range users {
		user.Password = userJSONs[index].Password
		fmt.Println(user)
		db.Create(&user)
	}

}
