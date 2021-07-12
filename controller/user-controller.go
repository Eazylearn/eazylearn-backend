package controller

import (
	"fmt"

	"github.com/Cumbercubie/api"
	"github.com/Cumbercubie/enum"
	"github.com/Cumbercubie/repo"
	"github.com/labstack/echo/v4"
)

func GetUserByFirstnameAction(c echo.Context) error {
	firstName := c.QueryParams().Get("firstName")
	// if len(firstName) == 0 {
	// 	api.Respond(c, &enum.APIResponse{
	// 		Status:  enum.APIStatus.Invalid,
	// 		Message: "firstName is empty",
	// 	})
	// 	return nil
	// }
	user, err := repo.GetUserByFirstname(firstName)
	if err != nil {
		api.Respond(c, &enum.APIResponse{
			Status:  enum.APIStatus.Error,
			Message: fmt.Sprintf("Error getting user: %s", err),
		})
		return nil
	}
	api.Respond(c, &enum.APIResponse{
		Status:  enum.APIStatus.Ok,
		Data:    user,
		Message: fmt.Sprintf("Lấy user thành công"),
	})
	return nil
}
func UserControllerGroup(g *echo.Group) error {
	g.GET("/", GetUserByFirstnameAction)
	return nil
}
