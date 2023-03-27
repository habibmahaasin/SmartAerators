package handler

import "github.com/gofiber/fiber/v2"

func (u *usersHandler) Holder(c *fiber.Ctx) error {
	return c.Render("pages/index", fiber.Map{
		"Title": "Hello, Note 👋!",
	}, layouts)
}
