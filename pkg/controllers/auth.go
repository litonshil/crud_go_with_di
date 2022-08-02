package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	consts "github.com/litonshil/crud_go_echo/pkg/const"
	"github.com/litonshil/crud_go_echo/pkg/domain"
	"github.com/litonshil/crud_go_echo/pkg/types"
)

type Auth struct {
	authSvc domain.IAuth
}

// NewAuthController will initialize the controllers
func NewAuthController(authSvc domain.IAuth) *Auth {
	ac := &Auth{
		authSvc: authSvc,
	}
	return ac
}

func (ur *Auth) Login(c echo.Context) error {
	var user = new(types.UserLoginReq)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, consts.BadRequest)
	}

	response, err := ur.authSvc.Login(c, user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, response)

}
