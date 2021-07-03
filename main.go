package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/tamnk74/todolist-mysql-go/config"
	"github.com/tamnk74/todolist-mysql-go/database"
	"github.com/tamnk74/todolist-mysql-go/middlewares"
	"github.com/tamnk74/todolist-mysql-go/router"
	"github.com/tamnk74/todolist-mysql-go/schedulers"
	"github.com/tamnk74/todolist-mysql-go/utils/queue"
	"github.com/tamnk74/todolist-mysql-go/utils/redis"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = database.Connect()
	if err != nil {
		panic("Failed to connect database")
	}

	app := fiber.New(middlewares.HandleApiError())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	router.Init(app)
	schedulers.Init()
	redis.Init()
	queue.Init()

	log.Info("Starting API server at port " + config.PORT)
	app.Listen(":" + config.PORT)
}
