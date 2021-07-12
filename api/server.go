package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Cumbercubie/enum"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GetContent(c echo.Context, template interface{}) error {
	return json.Unmarshal([]byte(GetContentText(c)), &template)
}

func Respond(context echo.Context, response *enum.APIResponse) error {
	switch response.Status {
	case enum.APIStatus.Ok:
		return context.JSON(http.StatusOK, response)
	case enum.APIStatus.Error:
		return context.JSON(http.StatusInternalServerError, response)
	case enum.APIStatus.Forbidden:
		return context.JSON(http.StatusForbidden, response)
	case enum.APIStatus.Invalid:
		return context.JSON(http.StatusBadRequest, response)
	case enum.APIStatus.NotFound:
		return context.JSON(http.StatusNotFound, response)
	case enum.APIStatus.Unauthorized:
		return context.JSON(http.StatusUnauthorized, response)
	case enum.APIStatus.Existed:
		return context.JSON(http.StatusConflict, response)
	}

	return context.JSON(http.StatusBadRequest, response)
}

// GetContentText ...
func GetContentText(c echo.Context) string {
	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}
	return string(bodyBytes)
}

type APIServer struct {
	Echo *echo.Echo
}

type Handler = func(e echo.Context) error

func InitServer() *APIServer {
	server := &APIServer{Echo: echo.New()}
	server.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	return server
}
func (server *APIServer) Start(port string) {
	server.Echo.Start(port)
}
func (server *APIServer) SetHandler(method *enum.MethodValue, path string, fn Handler) error {
	switch method.Value {
	case enum.APIMethod.GET.Value:
		server.Echo.GET(path, fn)
	case enum.APIMethod.POST.Value:
		server.Echo.POST(path, fn)
	case enum.APIMethod.PUT.Value:
		server.Echo.PUT(path, fn)
	case enum.APIMethod.DELETE.Value:
		server.Echo.DELETE(path, fn)
	}
	return nil
}

type ControllerFunc func(g *echo.Group) error

func (server *APIServer) SetGroup(group string, fn ControllerFunc) {
	g := server.Echo.Group(group)
	fn(g)
}
