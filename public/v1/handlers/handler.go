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
		"title":      "Beranda",
		"name":       name,
		"listDevice": ListDevice,
	}, layouts)
}

func (u *viewHandler) Login(c *fiber.Ctx) error {
	return c.Render("pages/login", fiber.Map{
		"title":   "Login",
		"message": "",
	}, layouts)
}

func (u *viewHandler) ListDevice(c *fiber.Ctx) error {
	sess, err := u.store.Get(c)
	if err != nil {
		panic(err)
	}

	name := sess.Get("name")
	ListDevice, err := u.deviceUsecase.GetAllDevices()
	if err != nil {
		panic(err)
	}

	return c.Render("pages/list_device", fiber.Map{
		"title":      "Daftar Perangkat",
		"name":       name,
		"listDevice": ListDevice,
	}, layouts)
}

func (u *viewHandler) Report(c *fiber.Ctx) error {
	sess, err := u.store.Get(c)
	if err != nil {
		panic(err)
	}

	name := sess.Get("name")
	History, err := u.deviceUsecase.GetDeviceHistory()
	if err != nil {
		panic(err)
	}

	return c.Render("pages/report", fiber.Map{
		"title":   "Laporan",
		"name":    name,
		"history": History,
	}, layouts)
}
