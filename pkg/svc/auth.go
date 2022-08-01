package svc

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	consts "github.com/litonshil/crud_go_echo/pkg/const"
	"github.com/litonshil/crud_go_echo/pkg/domain"
	"github.com/litonshil/crud_go_echo/pkg/models"
	"github.com/litonshil/crud_go_echo/pkg/token"
	"github.com/litonshil/crud_go_echo/pkg/types"
)

type IAuth interface {
	Login(c echo.Context, user *types.User) (*types.Token, error)
}

type auth struct {
	userRepo domain.IUsersRepo
}

func NewAuthService(userRepo domain.IUsersRepo) IAuth {
	return &auth{
		userRepo: userRepo,
	}
}

func mathcedCredentials(user *types.User, model_user *models.User) error {
	fmt.Println(user, model_user)
	if user.Password == model_user.Password && user.Type == model_user.Type {
		return nil
	}
	return errors.New("credintials not matched")
}

func (ur *auth) Login(c echo.Context, user *types.User) (*types.Token, error) {

	var model_user = new(models.User)
	var tokens = new(types.Token)

	model_user, err := ur.userRepo.GetUserByEmail(user.Email)
	credErr := mathcedCredentials(user, model_user)

	if model_user.Email == "" || err != nil || credErr != nil {
		return tokens, c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}

	token, refresh_token, err := token.GenerateUserTokens(model_user.Email, model_user.Id, model_user.Type)
	tokens.User_Token = token
	tokens.User_Refreshtoken = refresh_token

	if err != nil {
		return tokens, c.JSON(http.StatusInternalServerError, err.Error())
	}
	// return c.JSON(http.StatusOK, tokens)
	return tokens, nil
}
