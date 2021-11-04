package authentication

import (
	"time"
	"github.com/httpmon/user/config"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(id int, cfg config.JWT) (string, error) {
	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["github.com/httpmon/user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * time.Duration(cfg.Expiration)).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(cfg.Secret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidateToken(token string, cfg config.JWT) (in bool, i int) {
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.Secret), nil
	})
	if err != nil {
		return false, 0
	}

	auth := claims["authorized"].(bool)
	exp := claims["exp"].(float64)
	id := claims["github.com/httpmon/user_id"].(float64)

	if auth && exp > float64(time.Now().Unix()) {
		return true, int(id)
	}

	return false, int(id)
}
