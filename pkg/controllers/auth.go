package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	consts "github.com/litonshil/crud_go_echo/pkg/const"
	"github.com/litonshil/crud_go_echo/pkg/domain"
	"github.com/litonshil/crud_go_echo/pkg/types"
	"github.com/litonshil/crud_go_echo/pkg/utils"
)

type Auth struct {
	authSvc domain.IAuthSvc
}

// NewAuthController will initialize the controllers
func NewAuthController(authSvc domain.IAuthSvc) *Auth {
	ac := &Auth{
		authSvc: authSvc,
	}
	return ac
}

// Registration create a user
func (ur *Auth) Registration(c echo.Context) error {
	var user = new(types.UserRegisterReq)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, consts.BadRequest)
	}

	if validationerr := user.Validate(); validationerr != nil {
		return c.JSON(http.StatusInternalServerError, validationerr.Error())
	}

	if err := ur.authSvc.CreateUser(user); err != nil {
		if err.Error() == "email exist" {
			return c.JSON(http.StatusConflict, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Send username and password via email
	if err := utils.SendEmail(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "user created successfullys")
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
