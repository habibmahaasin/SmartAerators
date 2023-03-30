package middleware

import (
	"SmartAerators/infrastructures/config"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func AuthPages() func(c *fiber.Ctx) error {
	config, _ := config.New()
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Redirect("/login", fiber.StatusTemporaryRedirect)
		},
		TokenLookup:   "cookie:token",
		SigningMethod: "HS256",
		SigningKey:    []byte(config.App.Secret_key),
	})
}
