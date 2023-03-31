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
		return err
	}

	Antares_Device_Id := strings.Replace(webhookData.First.M2m_nev.M2m_rep.M2m_cin.Pi, "/antares-cse/cnt-", "", -1)
	_, err := dc.deviceUseCase.DataFromWebhook(webhookData.First.M2m_nev.M2m_rep.M2m_cin.Con, Antares_Device_Id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
