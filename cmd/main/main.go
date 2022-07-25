package main

import (
	"github.com/labstack/echo/v4"
	"github.com/litonshil/crud_go_echo/pkg/routes"
)

func main() {

	e := echo.New()

	routes.User(e)
	e.Start(":8080")

}
