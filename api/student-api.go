package api

import (
	"net/http"
	"strconv"

	"github.com/Cumbercubie/repo"
	"github.com/labstack/echo/v4"
)

func GetStudentById(c echo.Context) error {
	studentId, err := strconv.ParseInt(c.QueryParam("studentId"), 10, 0)
	if err != nil {
		panic(err.Error())
	} else {
		student := repo.GetStudentById(int32(studentId))
		return c.JSON(http.StatusOK, student)
	}
}
