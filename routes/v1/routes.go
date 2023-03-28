package routes

import (
	"SmartAerators/infrastructures/database"
	userController "SmartAerators/modules/v1/users/interfaces/controllers"
	userHandler "SmartAerators/public/v1/handler"

	"github.com/gofiber/fiber/v2"
)

func ParseTemplates(router *fiber.App) *fiber.App {
	router.Static("/", "./public")
	router.Static("/css", "./public/v1/assets/css")
	router.Static("/js", "./public/v1/assets/js")
	router.Static("/img", "./public/v1/assets/img")
	router.Static("/vendor", "./public/v1/assets/vendor")
	return router
}

func Routes(router *fiber.App, db database.Database) *fiber.App {
	userController := userController.UserController(db)
	userHandler := userHandler.Handler(db)

	views := router.Group("")
	views.Get("/", userHandler.Holder)

	api := router.Group("/api/v1")
	api.Get("/users", userController.Holder)

	router = ParseTemplates(router)
	return router
}
