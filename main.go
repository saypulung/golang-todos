package main

import (
	"fmt"

	"maspulung/gotodo/app/entities"
	"maspulung/gotodo/app/routes"
	"maspulung/gotodo/config"
	"maspulung/gotodo/config/database"
	"maspulung/gotodo/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	config.Init()
	database.Connect()
	database.Migrate(&entities.User{}, &entities.Todo{})

	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	fmt.Println("Engine start")

	// Konfigurasi CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173", // URL frontend Vite
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
		MaxAge:           300, // Durasi cache preflight request dalam detik
	}))

	app.Use(logger.New())

	routes.AuthRoutes(app)
	routes.TodoRoutes(app)

	fmt.Println("Routes registered")
	if err := app.Listen(fmt.Sprintf(":%v", config.PORT)); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
