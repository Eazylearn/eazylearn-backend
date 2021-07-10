package controller

import (
	"github.com/Cumbercubie/repo"
	"github.com/labstack/echo/v4"
)

func TestControllerGroup(g *echo.Group) error {
	g.GET("/test", repo.CreateTestAction)
	return nil
}
