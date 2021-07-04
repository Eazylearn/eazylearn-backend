package main

import (
	"net/http"

	"github.com/Cumbercubie/api"
	"github.com/labstack/echo/v4"
)

func initServer() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Wo!")
	})
	e.GET("/student", api.GetStudentById)
	e.Logger.Fatal(e.Start(":8081"))
}

func main() {
	initServer()
}
