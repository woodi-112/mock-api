package main

import (
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type MockAPI interface {
	Get(string, int, interface{})
	Post(string, int, interface{})
	Put(string, int, interface{})
	Patch(string, int, interface{})
	Delete(string, int, interface{})
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

// Post is a wrapper function that will add a new get handler to the server.
func (s *Server) Post(path string, responseCode int, responseBody interface{}) {
	s.echo.POST(path, func(c echo.Context) error {
		return c.JSON(responseCode, responseBody)
	})
}

// Put is a wrapper function that will add a new get handler to the server.
func (s *Server) Put(path string, responseCode int, responseBody interface{}) {
	s.echo.PUT(path, func(c echo.Context) error {
		return c.JSON(responseCode, responseBody)
	})
}

// Patch is a wrapper function that will add a new get handler to the server.
func (s *Server) Patch(path string, responseCode int, responseBody interface{}) {
	s.echo.PATCH(path, func(c echo.Context) error {
		return c.JSON(responseCode, responseBody)
	})
}

// Delete is a wrapper function that will add a new get handler to the server.
func (s *Server) Delete(path string, responseCode int, responseBody interface{}) {
	s.echo.DELETE(path, func(c echo.Context) error {
		return c.JSON(responseCode, responseBody)
	})
}
