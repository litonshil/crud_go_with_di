package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/litonshil/crud_go_echo/pkg/controllers"
)

func Auth(e *echo.Echo, cr *controllers.Auth) {

	e.POST("/users/login", cr.Login)

}
