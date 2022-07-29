package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/litonshil/crud_go_echo/pkg/domain"
	m "github.com/litonshil/crud_go_echo/pkg/middleware"
)

type users struct {
	urepo domain.IUsers
}

func User(e *echo.Echo, u domain.IUsers) {

	cr := &users{
		urepo: u,
	}

	sub := e.Group("/user", m.Authenticate)
	sub.POST("/registration", cr.urepo.Registration)
	sub.GET("/users", cr.urepo.GetAllUsers)
	sub.GET("/:id", cr.urepo.GetAUsers)
	sub.PUT("/:id", cr.urepo.UpdateUser)
	sub.DELETE("/:id", cr.urepo.DeleteUser)

}
