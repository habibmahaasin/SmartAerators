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

func (dc *DevicesController) Control(c *fiber.Ctx) error {
	id := c.Params("id")
	mode := c.Params("mode")
	antares_id := c.Params("antares")
	power := c.Params("power")
	token := "862b34fe2de548cc:cdf66d91b12db8d2"

	err := dc.deviceUseCase.Control(id, power, mode)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	postAntares := dc.deviceUseCase.PostControlAntares(antares_id, token, power, mode)
	if postAntares != nil {
		fmt.Println(postAntares)
		return nil
	}

	return c.Redirect("/daftar-perangkat")
}
