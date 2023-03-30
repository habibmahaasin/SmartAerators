package repositories

import (
	"SmartAerators/modules/v1/users/domain"
)

func (r *Repository) GetUserByEmail(email string) (domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).Find(&user).Error
	return user, err
}
