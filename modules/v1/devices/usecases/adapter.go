package usecases

import (
	"SmartAerators/modules/v1/devices/domain"
	repositoriesDevice "SmartAerators/modules/v1/devices/interfaces/repositories"
)

type DeviceUsecasePresenter interface {
	DataFromWebhook(sensorData string, antaresDeviceID string) (domain.SensorData, error)
	GetDeviceByAntares(antaresDeviceID string) (domain.Device, error)
	GetAllDevices() ([]domain.Device, error)
	Control(id string, power string, mode string) error
	PostControlAntares(antares_id string, token string, power string, mode string) error
	GetDeviceHistory() ([]domain.DeviceHistory, error)
}

type DeviceUsecase struct {
	repository repositoriesDevice.DeviceRepositoryPresenter
}

func NewDeviceUsecase(repositories repositoriesDevice.DeviceRepositoryPresenter) *DeviceUsecase {
	return &DeviceUsecase{repositories}
}
