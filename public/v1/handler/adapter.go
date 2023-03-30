package handler

import (
	userRepository "SmartAerators/modules/v1/users/interfaces/repositories"
	"SmartAerators/modules/v1/users/usecases"
	userUsecase "SmartAerators/modules/v1/users/usecases"

	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

var layouts string = "layouts/main"

type usersHandler struct {
	usersUsecase usecases.UsecasePresenter
	store        session.Store
}

func NewUsersHandler(usersUsecase usecases.UsecasePresenter, store session.Store) *usersHandler {
	return &usersHandler{usersUsecase, store}
}

func Handler(db *gorm.DB, store *session.Store) *usersHandler {
	repository := userRepository.NewRepository(db)
	usecase := userUsecase.NewUsecase(repository)

	return NewUsersHandler(usecase, *store)
}
