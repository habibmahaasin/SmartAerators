package controllers

import (
	database "SmartAerators/infrastructures/database"
	userRepository "SmartAerators/modules/v1/users/interfaces/repositories"
	userUsecase "SmartAerators/modules/v1/users/usecases"
	token "SmartAerators/package/jwt"
)

type UsersController struct {
	userUseCase userUsecase.UsecasePresenter
	jwtoken     token.JwToken
}

func NewUserController(userUsecase userUsecase.UsecasePresenter, jwtoken token.JwToken) *UsersController {
	return &UsersController{userUsecase, jwtoken}
}

func UserController(db database.Database) *UsersController {
	repository := userRepository.NewRepository(db)
	usecase := userUsecase.NewUsecase(repository)
	jwtoken := token.NewJwToken()

	return NewUserController(usecase, jwtoken)
}
