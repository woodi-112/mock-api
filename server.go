package main

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

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

	// Middleware
	s.echo.Use(middleware.Logger())
	s.echo.Use(middleware.Recover())

	// Routes
	s.echo.GET("/", hello)

	// Start server
	s.echo.Logger.Fatal(s.echo.Start(":8000"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
