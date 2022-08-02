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
	"github.com/litonshil/crud_go_echo/pkg/utils"
)

type auth struct {
	userRepo domain.IUsersRepo
}

func NewAuthService(userRepo domain.IUsersRepo) domain.IAuthSvc {
	return &auth{
		userRepo: userRepo,
	}
}

func mathcedCredentials(user *types.UserLoginReq, model_user *models.User) error {
	fmt.Println(user, model_user)
	if user.Password == model_user.Password && user.Type == model_user.Type {
		return nil
	}
	return errors.New("credintials not matched")
}

func (ur *auth) CreateUser(user *types.UserRegisterReq) error {
	var model_user = new(models.User)
	respErr := utils.StructToStruct(user, &model_user)
	if respErr != nil {
		return respErr
	}

	saveErr := ur.userRepo.CreateUser(model_user)
	if saveErr != nil {
		return saveErr
	}
	return saveErr
}

func (ur *auth) Login(c echo.Context, user *types.UserLoginReq) (*types.Token, error) {

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
