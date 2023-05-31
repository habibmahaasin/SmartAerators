package controllers

import (
	"SmartAerators/modules/v1/devices/domain"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (dc *DevicesController) SubscribeWebhook(c *fiber.Ctx) error {
	var webhookData domain.ObjectAntares1

	if err := c.BodyParser(&webhookData); err != nil {
		fmt.Println(err)
		return nil
	}

	Antares_Device_Id := strings.Replace(webhookData.First.M2m_nev.M2m_rep.M2m_cin.Pi, "/antares-cse/cnt-", "", -1)
	_, err := dc.deviceUseCase.DataFromWebhook(webhookData.First.M2m_nev.M2m_rep.M2m_cin.Con, Antares_Device_Id)
	if err != nil {
		fmt.Println("error : ", err)
		return nil
	}

	return nil
}

func (dc *DevicesController) PowerControl(c *fiber.Ctx) error {
	id := c.Params("id")
	power := c.Params("power")
	dc.deviceUseCase.PowerControl(id, power)

	return c.Redirect("/daftar-perangkat")
}

func (dc *DevicesController) ModeControl(c *fiber.Ctx) error {
	id := c.Params("id")
	mode := c.Params("mode")
	dc.deviceUseCase.ModeControl(id, mode)

	return c.Redirect("/daftar-perangkat")
}
