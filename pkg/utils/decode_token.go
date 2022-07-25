package utils

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/litonshil/crud_go_echo/config"
	"github.com/litonshil/crud_go_echo/pkg/types"
)

func DecodeToken(usertoken string) (*types.SignedUserDetails, error) {
	var claims = &types.SignedUserDetails{}
	_, err := jwt.ParseWithClaims(usertoken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().SecretKey), nil
	})
	return claims, err
}