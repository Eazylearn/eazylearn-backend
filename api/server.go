package api

import (
	"github.com/Cumbercubie/enum"
	"github.com/labstack/echo/v4"
)

type APIServer struct {
	Echo *echo.Echo
}

type Handler = func(e echo.Context) error

func InitServer() APIServer {
	server := APIServer{Echo: echo.New()}
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
