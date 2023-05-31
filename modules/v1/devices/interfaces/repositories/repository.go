package repositories

import (
	"SmartAerators/modules/v1/devices/domain"
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

	err := dr.db.Raw("select * from devices d inner join device_status ds ON d.status_id = ds.status_id inner join device_mode dm on d.mode_id = dm.mode_id").Scan(&device).Error
	return device, err
}

func (dr *DeviceRepository) PowerControl(id string, power string) error {
	err := dr.db.Exec("UPDATE devices SET status_id  = ?, date_updated = ? WHERE device_id = ?", power, time.Now(), id).Error
	return err
}

func (dr *DeviceRepository) ModeControl(id string, mode string) error {
	err := dr.db.Exec("UPDATE devices SET mode_id  = ?, date_updated = ? WHERE device_id = ?", mode, time.Now(), id).Error
	return err
}
