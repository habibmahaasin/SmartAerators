package usecases

import (
	"SmartAerators/modules/v1/users/domain"
	repositoriesUsers "SmartAerators/modules/v1/users/interfaces/repositories"
)

type UserUsecasePresenter interface {
	Holder() string
	Login(inputLogin domain.InputLogin) (domain.User, error)
}

type UserUsecase struct {
	repository repositoriesUsers.UserRepositoryPresenter
}

func NewUserUsecase(repositories repositoriesUsers.UserRepositoryPresenter) *UserUsecase {
	return &UserUsecase{repositories}
}
