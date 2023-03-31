package repositories

import (
	"SmartAerators/modules/v1/devices/domain"
	"fmt"
	"time"
)

func (dr *DeviceRepository) ExportDataWebhook(Device_id string, input domain.SensorData) error {
	err := dr.db.Exec("INSERT INTO device_history (device_id, status_id, data, history_date) VALUES (?,?,?,?)", Device_id, input.Status_device, input.Data, time.Now()).Error
	err = dr.db.Exec("UPDATE devices SET status_id = ?, date_updated = ? WHERE device_id = ?", input.Status_device, time.Now(), Device_id).Error
	return err
}

func (dr *DeviceRepository) GetDeviceByAntares(antaresDeviceID string) (domain.Device, error) {
	var device domain.Device
	err := dr.db.Where("antares_id = ?", antaresDeviceID).Find(&device).Error
	return device, err
}

func (dr *DeviceRepository) GetAllDevices() ([]domain.Device, error) {
	var device []domain.Device

	err := dr.db.Find(&device).Error
	fmt.Println(device)
	return device, err
}
