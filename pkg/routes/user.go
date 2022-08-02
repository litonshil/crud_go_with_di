package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/litonshil/crud_go_echo/pkg/controllers"
	// "github.com/litonshil/crud_go_echo/pkg/domain"
	// m "github.com/litonshil/crud_go_echo/pkg/middleware"
)

func InitRoute(e *echo.Echo, u *controllers.UserRepo, auth *controllers.Auth) {

	sub := e.Group("/user")

	sub.PUT("/:id", u.UpdateUser)
	sub.DELETE("/:id", u.DeleteUser)


	sub.POST("/registration", u.Registration)
	sub.GET("/users", u.GetUsers)
	sub.GET("/:id", u.GetUser)

	
	e.POST("/users/login", auth.Login)

}
