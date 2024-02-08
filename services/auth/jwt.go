package auth

import (
	"errors"
	"time"

	"github.com/bloiseleo/chorumecon/env"
	"github.com/golang-jwt/jwt/v5"
)

func getExpirationTime() int64 {
	if env.Env("GIN_MODE", "debug") == "debug" {
		return time.Now().Add(time.Second * 120).Unix()
	}
	return time.Now().Add(time.Second * 10).Unix()
}

func CreateToken(id string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": id,
			"exp": getExpirationTime(),
		})
	tokenString, err := token.SignedString([]byte(env.Env("secret", "secret")))
	if err != nil {
		panic(err)
	}
	return tokenString
}

func ValidateJWT(token string) (*jwt.Token, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(env.Env("secret", "secret")), nil
	})
	if err != nil {
		return nil, err
	}
	if !t.Valid {
		return nil, errors.New("Expired")
	}
	return t, nil
}

func TokenExpired(err error) bool {
	return errors.Is(err, jwt.ErrTokenExpired)
}
