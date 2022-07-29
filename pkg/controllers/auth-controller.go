package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	consts "github.com/litonshil/crud_go_echo/pkg/const"
	"github.com/litonshil/crud_go_echo/pkg/domain"

	// "github.com/litonshil/crud_go_echo/pkg/models"
	// "github.com/litonshil/crud_go_echo/pkg/svc"
	svcImpl "github.com/litonshil/crud_go_echo/pkg/svc/impl"
	"github.com/litonshil/crud_go_echo/pkg/types"
)

type auth struct {
	authSvc svcImpl.IAuth
}

// NewAuthController will initialize the controllers
func NewAuthController(authSvc svcImpl.IAuth) domain.IAuthController {
	ac := &auth{
		authSvc: authSvc,
		// uSvc:    uSvc,
	}
	return ac
}

func (ur *auth) Login(c echo.Context) error {
	var user = new(types.User)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, consts.BadRequest)
	}

	response, err := ur.authSvc.Login(c, user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, response)

}
