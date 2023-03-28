package main

import (
	"SmartAerators/infrastructures/config"
	"SmartAerators/infrastructures/database"
	"SmartAerators/routes/v1"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
)

func Define() (*fiber.App, database.Database) {
	database := database.NewDatabase()
	engine := html.New("./public/v1/templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			err = ctx.Status(code).Render(fmt.Sprintf("errors/%d", code), fiber.Map{
				"Title": "Error " + fmt.Sprintf("%d", code),
			})

			if err != nil {
				return ctx.Status(500).Render("errors/500", fiber.Map{
					"Title": fiber.ErrInternalServerError.Message,
				})
			}

			return nil
		},
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
