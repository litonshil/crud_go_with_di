package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/litonshil/crud_go_echo/pkg/controllers"
	m "github.com/litonshil/crud_go_echo/pkg/middleware"
)

func InitRoute(e *echo.Echo, u *controllers.UserRepo, auth *controllers.Auth) {

	sub := e.Group("/user")

	// public
	sub.GET("/users", u.GetUsers)
	sub.GET("/:id", u.GetUser)
	sub.POST("/registration", auth.Registration)
	sub.POST("/login", auth.Login)

	//private
	sub.PUT("/:id", u.UpdateUser, m.Authenticate)
	sub.DELETE("/:id", u.DeleteUser, m.Authenticate)

}
