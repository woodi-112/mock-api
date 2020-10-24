package main

import (
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type MockAPI interface {
	Get(string, int, interface{})
}

type Server struct {
	echo *echo.Echo
}

func New() *Server {
	echo := echo.New()

	return &Server{
		echo: echo,
	}
}

func (s *Server) Start() {
	s.echo.Use(middleware.Logger())
	s.echo.Use(middleware.Recover())
	s.echo.Logger.Fatal(s.echo.Start(":8000"))
}

// Get is a wrapper function that will add a new get handler to the server.
func (s *Server) Get(path string, responseCode int, responseBody interface{}) {
	s.echo.GET(path, func(c echo.Context) error {
		return c.JSON(responseCode, responseBody)
	})
}
