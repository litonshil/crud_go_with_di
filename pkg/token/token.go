package token

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/litonshil/crud_go_echo/config"
	"github.com/litonshil/crud_go_echo/pkg/types"
)

func GenerateUserTokens(email string, id int, usertype string) (signed_token string, signed_refreshtoken string, err error) {

	claims := &types.SignedUserDetails{
		Email: email,
		Id:    id,
		Type:  usertype,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshclaims := &types.SignedUserDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.GetConfig().SecretKey))

	if err != nil {
		return token, "", err
	}
	refresh_token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshclaims).SignedString([]byte(config.GetConfig().SecretKey))

	if err != nil {
		return token, refresh_token, err
	}

	return token, refresh_token, nil
}
