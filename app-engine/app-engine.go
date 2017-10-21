package main

import (
	"net/http"

	"github.com/labstack/echo"
)

var e = createMux()

func createMux() *echo.Echo {
	e := echo.New()

	http.Handle("/", e)
	return e
}
