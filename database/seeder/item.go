package seeder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tamnk74/todolist-mysql-go/models"
	"gorm.io/gorm"
)

// Migration for table items, models item
func SeedItem(db *gorm.DB) {
	jsonFile, err := os.Open("database/data/items.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("opened items.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var items []models.Item
	json.Unmarshal([]byte(byteValue), &items)
	db.Create(&items)
}
