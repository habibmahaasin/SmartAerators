package controllers

import (
	"SmartAerators/modules/v1/users/domain"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (uc *UsersController) Login(c *fiber.Ctx) error {
	var inputLogin domain.InputLogin

	if err := c.BodyParser(&inputLogin); err != nil {
		return c.SendString("Error : " + err.Error())
	}

	user, err := uc.userUseCase.Login(inputLogin)
	if err != nil {
		return c.Render("pages/login", fiber.Map{
			"title":   "Login",
			"message": err.Error(),
		}, layouts)
	}

	token, err := uc.jwtoken.GenerateToken(user.User_id, user.Full_name, user.Role_id)
	c.Cookie(&fiber.Cookie{
		Name:  "token",
		Value: token,
	})

	sess, err := uc.store.Get(c)
	if err != nil {
		panic(err)
	}

	sess.Set("name", user.Full_name)
	sess.Set("user_id", user.User_id)
	sess.Set("role", user.Role_id)
	sess.SetExpiry(time.Minute * 15) //temporary
	if err := sess.Save(); err != nil {
		panic(err)
	}

	return c.Redirect("/")
}
