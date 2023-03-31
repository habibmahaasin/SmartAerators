package controllers

import (
	deviceRepository "SmartAerators/modules/v1/devices/interfaces/repositories"
	deviceUsecase "SmartAerators/modules/v1/devices/usecases"

	"gorm.io/gorm"
)

type DevicesController struct {
	deviceUseCase deviceUsecase.DeviceUsecasePresenter
}

func NewDeviceController(deviceUsecase deviceUsecase.DeviceUsecasePresenter) *DevicesController {
	return &DevicesController{deviceUsecase}
}

func DeviceController(db *gorm.DB) *DevicesController {
	repository := deviceRepository.NewDeviceRepository(db)
	usecase := deviceUsecase.NewDeviceUsecase(repository)

	return NewDeviceController(usecase)
}
