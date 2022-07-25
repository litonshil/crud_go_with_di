package helpers

import (
	jwt "github.com/dgrijalva/jwt-go"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/litonshil/crud_go_echo/config"
	"github.com/litonshil/crud_go_echo/pkg/types"
)

func VerifyToken(usertoken string) (bool, error) {
	claims := &types.SignedUserDetails{}
	flag := false
	token, err := jwt.ParseWithClaims(usertoken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().SecretKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {

			return flag, err
		}

		return flag, err
	}
	if !token.Valid {

		return flag, err
	}
	flag = true

	return flag, nil

}
