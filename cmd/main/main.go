package main

import (
	"github.com/labstack/echo/v4"
	"github.com/litonshil/crud_go_echo/pkg/routes"
)

// func start(e *echo.Echo) {
// 	routes.User(e)

// }

func main() {

	e := echo.New()

	// start(e)

	routes.User(e)
	e.Start(":8080")

}
