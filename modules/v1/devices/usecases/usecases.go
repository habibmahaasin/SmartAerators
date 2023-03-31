package usecases

import (
	"SmartAerators/modules/v1/devices/domain"
	"encoding/json"
	"fmt"
)

func (du *DeviceUsecase) DataFromWebhook(sensorData string, antaresDeviceID string) (domain.SensorData, error) {
	var data domain.SensorData

	err := json.Unmarshal([]byte(sensorData), &data)
	if err != nil {
		fmt.Println(err)
		return data, err
	}

	getDetailDevice, err := du.GetDeviceByAntares(antaresDeviceID)
	data.Device_id = getDetailDevice.Device_id
	if data.Status_device != 11 {
		data.Status_device = 10
	}

	// getDeviceDetail := du.
	err = du.repository.ExportDataWebhook(data.Device_id, data)
	return data, err
}

func (du *DeviceUsecase) GetDeviceByAntares(antaresDeviceID string) (domain.Device, error) {
	return du.repository.GetDeviceByAntares(antaresDeviceID)
}

func (du *DeviceUsecase) GetAllDevices() ([]domain.Device, error) {
	return du.repository.GetAllDevices()
}
