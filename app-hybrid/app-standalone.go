// +build !appengine

package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return e
}

func main() {
	e.Logger.Fatal(e.Start(":8080"))
}
