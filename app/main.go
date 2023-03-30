package main

import (
	"SmartAerators/infrastructures/config"
	"SmartAerators/infrastructures/database"
	"SmartAerators/routes/v1"
	"errors"

	errorhandling "SmartAerators/package/error_handling"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
)

func Define() *fiber.App {
	engine := html.New("./public/v1/templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			err = errorhandling.PageNotFound(ctx, code)
			if err != nil {
				return errorhandling.InternalErrors(ctx)
			}

			return nil
		},
	})

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	return app
}

func main() {
	router := Define()
	db := database.Init()
	config, _ := config.New()
	router = routes.Routes(router, db)

	router.Listen(config.App.Url + ":" + config.App.Port)
}
