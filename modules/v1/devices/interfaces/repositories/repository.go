package repositories

import (
	"SmartAerators/modules/v1/devices/domain"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func (dr *DeviceRepository) ExportDataWebhook(Device_id string, input domain.SensorData) error {
	err := dr.db.Exec("INSERT INTO device_history (status_id, mode_id, device_id, temperature, ph, dissolved_oxygen, history_date) VALUES (?,?,?,?,?,?,?)", input.Status_device, input.Device_mode, Device_id, input.Temperature, input.Ph, input.Dissolved_oxygen, time.Now()).Error
	err = dr.db.Exec("UPDATE devices SET status_id  = ?, mode_id  = ?, date_updated = ? WHERE device_id = ?", input.Status_device, input.Device_mode, time.Now(), Device_id).Error
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

func (dr *DeviceRepository) Control(id string, power string, mode string) error {
	err := dr.db.Exec("UPDATE devices SET status_id  = ?, mode_id  = ?, date_updated = ? WHERE device_id = ?", power, mode, time.Now(), id).Error
	return err
}

func (dr *DeviceRepository) PostControlAntares(antares_id string, token string, power string, mode string) error {
	data := "\r\n{\r\n  \"m2m:cin\": {\r\n    \"con\": \"{ \\\"aeratorMode\\\":" + mode + ", \\\"statusDevice\\\":" + power + "}\"\r\n    }\r\n}"

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("POST", "https://platform.antares.id:8443/~/antares-cse/cnt-"+antares_id, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return err
	}

	req.Header.Set("X-M2M-Origin", token)
	req.Header.Set("Content-Type", "application/json;ty=4")
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (dr *DeviceRepository) GetDeviceHistory() ([]domain.DeviceHistory, error) {
	var DeviceHistory []domain.DeviceHistory
	fmt.Println(DeviceHistory)
	err := dr.db.Raw("select d.device_name, ds.status_name, dm.mode_name, dh.ph, dh.temperature, dh.dissolved_oxygen, dh.history_date from  device_history dh inner join devices d on dh.device_id = d.device_id inner join device_status ds on dh.status_id = ds.status_id inner join device_mode dm on dh.mode_id = dm.mode_id ORDER BY dh.history_id DESC").Scan(&DeviceHistory).Error
	return DeviceHistory, err
}
