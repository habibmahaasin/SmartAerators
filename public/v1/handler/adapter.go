package handler

import (
	database "SmartAerators/infrastructures/database"
	userRepository "SmartAerators/modules/v1/users/interfaces/repositories"
	"SmartAerators/modules/v1/users/usecases"
	userUsecase "SmartAerators/modules/v1/users/usecases"
)

var layouts string = "layouts/main"

type usersHandler struct {
	usersUsecase usecases.UsecasePresenter
}

func NewUsersHandler(usersUsecase usecases.UsecasePresenter) *usersHandler {
	return &usersHandler{usersUsecase}
}

func Handler(db database.Database) *usersHandler {
	repository := userRepository.NewRepository(db)
	usecase := userUsecase.NewUsecase(repository)

	return NewUsersHandler(usecase)
}
