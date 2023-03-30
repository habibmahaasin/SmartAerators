package usecases

import (
	"SmartAerators/modules/v1/users/domain"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func (u *Usecase) Holder() string {
	a := "hello"
	return a
}

func (u *Usecase) Login(inputLogin domain.InputLogin) (domain.User, error) {
	email := inputLogin.Email
	password := inputLogin.Password

	findUser, _ := u.repository.GetUserByEmail(email)
	if findUser.User_id == "" {
		return findUser, errors.New("Email Doesnt Exist")
	}

	comparePassword := bcrypt.CompareHashAndPassword([]byte(findUser.Password), []byte(password))
	if comparePassword != nil {
		return findUser, errors.New("Invalid Email or Password")
	}

	return findUser, nil
}
