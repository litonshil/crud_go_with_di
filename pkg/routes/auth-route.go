package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/litonshil/crud_go_echo/pkg/domain"
)

type userA struct {
	urepo domain.IAuth
}

func Auth(e *echo.Echo, u domain.IAuth) {

	cr := &userA{
		urepo: u,
	}
	e.POST("/users/login", cr.urepo.Login)

}
