package jwt

import (
	"SmartAerators/infrastructures/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwToken interface {
	GenerateToken(userID string, fullName string, role int) (string, error)
	ValidateToken(encode string) (*jwt.Token, error)
}

type jwtoken struct{}

func NewJwToken() *jwtoken {
	return &jwtoken{}
}

func (*jwtoken) GenerateToken(userID string, fullName string, role int) (string, error) {
	conf, _ := config.New()
	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	claim["full_name"] = fullName
	claim["role_id"] = role
	claim["exp"] = time.Now().Add(time.Minute * 15).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, _ := token.SignedString([]byte(conf.App.Secret_key))

	return signedToken, nil
}

func (*jwtoken) ValidateToken(encode string) (*jwt.Token, error) {
	conf, _ := config.New()
	token, err := jwt.Parse(encode, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(conf.App.Secret_key), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
