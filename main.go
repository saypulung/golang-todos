package main

import (
	"fmt"

	"maspulung/gotodo/app/dal"
	"maspulung/gotodo/app/routes"
	"maspulung/gotodo/config"
	"maspulung/gotodo/config/database"
	"maspulung/gotodo/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	config.Init()
	database.Connect()
	database.Migrate(&dal.User{}, &dal.Todo{})

	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	fmt.Println("Engine start")

	app.Use(logger.New())

	routes.AuthRoutes(app)
	routes.TodoRoutes(app)

	fmt.Println("Routes registered")
	if err := app.Listen(fmt.Sprintf(":%v", config.PORT)); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
