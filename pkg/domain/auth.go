package domain

import (
	"github.com/labstack/echo/v4"
	"github.com/litonshil/crud_go_echo/pkg/types"
)

type IAuthSvc interface {
	Login(c echo.Context, user *types.UserLoginReq) (*types.Token, error)
	CreateUser(user *types.UserRegisterReq) error
}
