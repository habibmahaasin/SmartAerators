package repositories

import (
	"SmartAerators/modules/v1/devices/domain"

	"gorm.io/gorm"
)

type DeviceRepositoryPresenter interface {
	ExportDataWebhook(Device_id string, input domain.SensorData) error
	GetDeviceByAntares(antaresDeviceID string) (domain.Device, error)
	GetAllDevices() ([]domain.Device, error)
	PowerControl(id string, power string) error
	ModeControl(id string, mode string) error
}

type DeviceRepository struct {
	db *gorm.DB
}

func NewDeviceRepository(db *gorm.DB) *DeviceRepository {
	return &DeviceRepository{db}
}
