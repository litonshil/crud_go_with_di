package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/litonshil/crud_go_echo/pkg/controllers"
	// "github.com/litonshil/crud_go_echo/pkg/domain"
	// m "github.com/litonshil/crud_go_echo/pkg/middleware"
)

// type users struct {
// 	urepo domain.IUsersController
// }

func User(e *echo.Echo, u *controllers.UserRepo) {

	sub := e.Group("/user")
	sub.POST("/registration", u.Registration)
	sub.GET("/users", u.GetAllUsers)
	sub.GET("/:id", u.GetAUsers)
	sub.PUT("/:id", u.UpdateUser)
	sub.DELETE("/:id", u.DeleteUser)

}
