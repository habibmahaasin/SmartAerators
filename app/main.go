package main

import (
	"SmartAerators/infrastructures/config"
	"SmartAerators/infrastructures/database"
	"SmartAerators/routes/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
)

func Define() (*fiber.App, database.Database) {
	database := database.NewDatabase()
	engine := html.New("./public/v1/templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	return app, database
}

func main() {
	router, db := Define()
	config, _ := config.New()
	router = routes.Routes(router, db)

	router.Listen(config.App.Url + ":" + config.App.Port)
}
