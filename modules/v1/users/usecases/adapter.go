package usecases

import (
	"SmartAerators/modules/v1/users/domain"
	repositoriesUsers "SmartAerators/modules/v1/users/interfaces/repositories"
)

type UsecasePresenter interface {
	Holder() string
	Login(inputLogin domain.InputLogin) (domain.User, error)
}

type Usecase struct {
	repository repositoriesUsers.RepositoryPresenter
}

func NewUsecase(repositories repositoriesUsers.RepositoryPresenter) *Usecase {
	return &Usecase{repositories}
}
