package errorhandling

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func InternalErrors(c *fiber.Ctx) error {
	return c.Render("errors/500", fiber.Map{
		"Title": "Internal Server Error",
		"Desc":  "Error 500",
	})
}

func PageNotFound(c *fiber.Ctx, code int) error {
	return c.Status(code).Render(fmt.Sprintf("errors/%d", code), fiber.Map{
		"Title": "Page Not Found",
		"Desc":  "Error " + fmt.Sprintf("%d", code),
	})
}
