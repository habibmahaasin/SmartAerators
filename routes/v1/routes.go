package routes

import (
	userController "SmartAerators/modules/v1/users/interfaces/controllers"
	userHandler "SmartAerators/public/v1/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

func ParseTemplates(router *fiber.App) *fiber.App {
	router.Static("/", "./public")
	router.Static("/css", "./public/v1/assets/css")
	router.Static("/js", "./public/v1/assets/js")
	router.Static("/img", "./public/v1/assets/img")
	router.Static("/vendor", "./public/v1/assets/vendor")
	return router
}

func Routes(router *fiber.App, db *gorm.DB) *fiber.App {
	store := session.New()
	userController := userController.UserController(db, store)
	userHandler := userHandler.Handler(db, store)

	router = ParseTemplates(router)

	api := router.Group("/api/v1")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to SmartAerators API")
	})

	router.Get("/login", userHandler.Login)
	router.Post("/login", userController.Login)

	pages := router.Group("")
	pages.Get("/", userHandler.Index)

	return router
}
