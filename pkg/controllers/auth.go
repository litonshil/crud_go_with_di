package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
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
		return c.JSON(http.StatusBadRequest, utils.RequestBodyParseErrorResponseMsg())
	}

	if validationerr := user.Validate(); validationerr != nil {
		return c.JSON(http.StatusBadRequest, utils.ValidationErrorMsg())
	}

	if err := ur.authSvc.CreateUser(user); err != nil {
		if err.Error() == "email exist" {
			return c.JSON(http.StatusConflict, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, utils.EntityCreationFailedMsg("User"))

	}

	// Send username and password via email
	if err := utils.SendEmail(user); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.MailSendingFailedMsg("User registration"))

	}

	return c.JSON(http.StatusCreated, "user created successfullys")
}

func (ur *Auth) Login(c echo.Context) error {
	var user = new(types.UserLoginReq)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, utils.RequestBodyParseErrorResponseMsg())
	}

	response, err := ur.authSvc.Login(c, user)

	if err != nil {
		switch err {
		case utils.ErrInvalidEmail, utils.ErrInvalidPassword:
			return c.JSON(http.StatusUnauthorized, utils.InvalidUserPassMsg())
		case utils.ErrCreateJwt:
			return c.JSON(http.StatusInternalServerError, utils.JwtCreateErrorMsg())
		default:
			return c.JSON(http.StatusInternalServerError, utils.SomethingWentWrongMsg())
		}
	}

	return c.JSON(http.StatusOK, response)

}
