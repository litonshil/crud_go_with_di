package domain

import (
	"github.com/labstack/echo/v4"
	"github.com/litonshil/crud_go_echo/pkg/types"
)

// type IAuthController interface {
// 	Login(c echo.Context) error
// }

type IAuth interface {
	Login(c echo.Context, user *types.UserLoginReq) (*types.Token, error)
}
