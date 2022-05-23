package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func helloController(c echo.Context) error {
	message := "Hello World from Golang"

	return c.JSON(http.StatusOK, message)
}

func main() {
	e := echo.New()

	e.GET("/", helloController)

	e.Start(":8000")
}
