package repo

import (
	"net/http"

	"github.com/Cumbercubie/model"
	"github.com/labstack/echo/v4"
)

func CreateTestAction(c echo.Context) error {
	c.JSON(http.StatusOK, model.CreateTest())
	return nil
}
