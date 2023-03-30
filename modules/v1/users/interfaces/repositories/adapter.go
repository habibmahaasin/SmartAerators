package repositories

import (
	"SmartAerators/modules/v1/users/domain"

	"gorm.io/gorm"
)

type RepositoryPresenter interface {
	GetUserByEmail(email string) (domain.User, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}
