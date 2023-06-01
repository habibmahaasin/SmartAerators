package domain

import "time"

type Device struct {
	Device_id       string
	Antares_id      string
	Device_name     string
	Device_location string
	Device_label    string
	Status_id       string
	Status_name     string
	Mode_id         string
	Mode_name       string
	Brand           string
	User_id         string
	Latitude        string
	Longitude       string
	Date_created    time.Time
	Date_updated    time.Time
}

type DeviceHistory struct {
	History_device_name      string    `gorm:"column:device_name"`
	History_status_name      string    `gorm:"column:status_name"`
	History_mode_name        string    `gorm:"column:mode_name"`
	History_ph               float32   `gorm:"column:ph"`
	History_temperature      float32   `gorm:"column:temperature"`
	History_dissolved_oxygen float32   `gorm:"column:dissolved_oxygen"`
	History_date             time.Time `gorm:"column:history_date"`
}
