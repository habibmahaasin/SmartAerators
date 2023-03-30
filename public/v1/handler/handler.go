package handler

import (
	"github.com/gofiber/fiber/v2"
)

func (u *usersHandler) Index(c *fiber.Ctx) error {
	sess, err := u.store.Get(c)
	if err != nil {
		panic(err)
	}

	name := sess.Get("name")

	return c.Render("pages/index", fiber.Map{
		"title": "Hello, Index 👋!",
		"name":  name,
	}, layouts)
}

func (u *usersHandler) Login(c *fiber.Ctx) error {
	return c.Render("pages/login", fiber.Map{
		"title": "Login",
	}, layouts)
}
