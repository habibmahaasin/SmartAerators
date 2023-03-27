package usecases

import repositoriesUsers "SmartAerators/modules/v1/users/interfaces/repositories"

type UsecasePresenter interface {
	Holder() string
}

type Usecase struct {
	repository repositoriesUsers.RepositoryPresenter
}

func NewUsecase(repositories repositoriesUsers.RepositoryPresenter) *Usecase {
	return &Usecase{repositories}
}
