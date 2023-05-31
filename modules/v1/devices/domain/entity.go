package domain

import "time"

type Device struct {
	Device_id       string
	Antares_id      string
	Device_name     string
	Device_location string
	Device_label    string
	Status_name     string
	Mode_name       string
	Brand           string
	User_id         string
	Latitude        string
	Longitude       string
	Date_created    time.Time
	Date_updated    time.Time
}
