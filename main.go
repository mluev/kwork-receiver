package main

import (
	"kworker/clients"
	"kworker/config"
	"kworker/handlers"
	"kworker/services"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	bot := clients.Init()

	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		services.SendNewTasks(bot)

		for range ticker.C {
			services.SendNewTasks(bot)
		}
	}()

	handlers.Init(bot)
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))


	log.Fatal(app.Listen(":" + config.Config("PORT")))
}
