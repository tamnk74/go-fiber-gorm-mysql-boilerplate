package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/tamnk74/todolist-mysql-go/config"
	"github.com/tamnk74/todolist-mysql-go/database"
	"github.com/tamnk74/todolist-mysql-go/middlewares"
	"github.com/tamnk74/todolist-mysql-go/router"
	"github.com/tamnk74/todolist-mysql-go/schedulers"
	"github.com/tamnk74/todolist-mysql-go/utils/queue"
	"github.com/tamnk74/todolist-mysql-go/utils/redis"
)

var db *gorm.DB
var err error

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = database.Connect()
	if err != nil {
		panic("Failed to connect database")
	}

	log.Info("Starting API server at port " + config.PORT)
	app := fiber.New(middlewares.HandleApiError())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	router.Init(app)
	schedulers.Init()
	redis.Init()
	queue.Init()
	app.Listen(":" + config.PORT)
}
