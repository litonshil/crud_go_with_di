package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/litonshil/crud_go_echo/pkg/controllers"
	"github.com/litonshil/crud_go_echo/pkg/middleware"
)

func User(e *echo.Echo) {
	sub := e.Group("/user",middleware.Authenticate)
	sub.POST("/registration", controllers.Registration)
	sub.POST("/login", controllers.Login)
	sub.GET("/users", controllers.GetAllUsers)
	sub.GET("/:id", controllers.GetAUsers)
	sub.PUT("/:id", controllers.UpdateUser)
	sub.DELETE("/:id", controllers.DeleteUser)

}
