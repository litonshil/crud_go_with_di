package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/litonshil/crud_go_echo/pkg/domain"
)

type userA struct {
	urepo domain.IAuthController
}

func Auth(e *echo.Echo, u domain.IAuthController) {

	cr := &userA{
		urepo: u,
	}
	e.POST("/users/login", cr.urepo.Login)

}