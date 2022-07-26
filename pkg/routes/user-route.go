package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/litonshil/crud_go_echo/pkg/controllers"
	"github.com/litonshil/crud_go_echo/pkg/database"

	// "github.com/litonshil/crud_go_echo/pkg/middleware"
	"github.com/litonshil/crud_go_echo/pkg/repository"
)

func User(e *echo.Echo) {

	db := database.GetDB()

	dbc := repository.NewDb(db)
	userc := controllers.NewUserController(dbc)

	sub := e.Group("/user")
	sub.POST("/registration", userc.Registration)
	sub.POST("/login", userc.Login)
	sub.GET("/users", userc.GetAllUsers)
	sub.GET("/:id", userc.GetAUsers)
	sub.PUT("/:id", userc.UpdateUser)
	sub.DELETE("/:id", userc.DeleteUser)

}
