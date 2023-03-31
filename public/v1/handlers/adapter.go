package handler

import (
	deviceRepository "SmartAerators/modules/v1/devices/interfaces/repositories"
	userRepository "SmartAerators/modules/v1/users/interfaces/repositories"

	deviceUsecase "SmartAerators/modules/v1/devices/usecases"
	userUsecase "SmartAerators/modules/v1/users/usecases"

	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

var layouts string = "layouts/main"

type viewHandler struct {
	usersUsecase  userUsecase.UserUsecasePresenter
	deviceUsecase deviceUsecase.DeviceUsecasePresenter
	store         session.Store
}

func NewViewHandler(usersUsecase userUsecase.UserUsecasePresenter, deviceUsecase deviceUsecase.DeviceUsecasePresenter, store session.Store) *viewHandler {
	return &viewHandler{usersUsecase, deviceUsecase, store}
}

func Handler(db *gorm.DB, store *session.Store) *viewHandler {
	userRepository := userRepository.NewUserRepository(db)
	userUsecase := userUsecase.NewUserUsecase(userRepository)

	deviceRepository := deviceRepository.NewDeviceRepository(db)
	deviceUsecase := deviceUsecase.NewDeviceUsecase(deviceRepository)

	return NewViewHandler(userUsecase, deviceUsecase, *store)
}
