package controllers

import "github.com/gofiber/fiber/v2"

func (uc *UsersController) Holder(c *fiber.Ctx) error {
	getHolderFromService := uc.userUseCase.Holder()
	getToken, err := uc.jwtoken.GenerateToken(1)

	if err != nil {
		return c.SendString("Error Generate Token : " + err.Error())
	}

	return c.SendString("Hasil Generate Service : " + getHolderFromService + "\nToken : " + getToken)
}
