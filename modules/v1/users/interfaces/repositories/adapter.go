package repositories

import (
	"SmartAerators/modules/v1/users/domain"

	"gorm.io/gorm"
)

type UserRepositoryPresenter interface {
	GetUserByEmail(email string) (domain.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}
