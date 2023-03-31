package handler

import (
	"github.com/gofiber/fiber/v2"
)

func (u *viewHandler) Index(c *fiber.Ctx) error {
	sess, err := u.store.Get(c)
	if err != nil {
		panic(err)
	}

	name := sess.Get("name")
	ListDevice, err := u.deviceUsecase.GetAllDevices()
	if err != nil {
		panic(err)
	}

	return c.Render("pages/index", fiber.Map{
		"title":      "Hello, Index 👋!",
		"name":       name,
		"listDevice": ListDevice,
	}, layouts)
}

func (u *viewHandler) Login(c *fiber.Ctx) error {
	return c.Render("pages/login", fiber.Map{
		"title": "Login",
	}, layouts)
}
