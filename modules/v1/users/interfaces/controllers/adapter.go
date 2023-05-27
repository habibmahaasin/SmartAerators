package controllers

import (
	userRepository "SmartAerators/modules/v1/users/interfaces/repositories"
	userUsecase "SmartAerators/modules/v1/users/usecases"
	token "SmartAerators/package/jwt"

	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

var layouts string = "layouts/main"

type UsersController struct {
	userUseCase userUsecase.UserUsecasePresenter
	jwtoken     token.JwToken
	store       session.Store
}

func NewUserController(userUsecase userUsecase.UserUsecasePresenter, jwtoken token.JwToken, store session.Store) *UsersController {
	return &UsersController{userUsecase, jwtoken, store}
}

func UserController(db *gorm.DB, store *session.Store) *UsersController {
	repository := userRepository.NewUserRepository(db)
	usecase := userUsecase.NewUserUsecase(repository)
	jwtoken := token.NewJwToken()

	return NewUserController(usecase, jwtoken, *store)
}
