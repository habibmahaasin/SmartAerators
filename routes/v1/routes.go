package routes

import (
	mid "SmartAerators/infrastructures/middleware"
	deviceController "SmartAerators/modules/v1/devices/interfaces/controllers"
	userController "SmartAerators/modules/v1/users/interfaces/controllers"
	viewHandler "SmartAerators/public/v1/handlers"

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
	deviceController := deviceController.DeviceController(db)
	viewHandler := viewHandler.Handler(db, store)

	router = ParseTemplates(router)

	api := router.Group("/api/v1")
	api.Post("/webhook", deviceController.SubscribeWebhook)

	router.Get("/login", viewHandler.Login)
	router.Post("/login", userController.Login)

	pages := router.Group("", mid.AuthPages())
	pages.Get("/", viewHandler.Index)
	pages.Get("/daftar-perangkat", viewHandler.ListDevice)
	pages.Get("/control/:id/:antares/:mode/:power", deviceController.Control)
	pages.Get("/logout", userController.Logout)

	return router
}
