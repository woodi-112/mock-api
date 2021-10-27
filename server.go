package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type MockAPI interface {
	Get(string, Response)
	Post(string, Response)
	Put(string, Response)
	Patch(string, Response)
	Delete(string, Response)
}

// Response struct to make things easier
type Response struct {
	Code int
	Body interface{}
}

// Server is he top level struct
type Server struct {
	echo *echo.Echo
	port int
}

// New creates and returns a new Server
func New(port int) *Server {
	echo := echo.New()

	return &Server{
		echo: echo,
		port: port,
	}
}

// Add default server
func NewDefault() *Server {
	echo := echo.New()

	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())

	return &Server{
		echo: echo,
		port: 8000,
	}
}

func (s *Server) Use(middleware echo.MiddlewareFunc) {
	s.echo.Use(middleware)
}

// Start the server with logger
func (s *Server) Start() {
	s.echo.Logger.Fatal(
		s.echo.Start(fmt.Sprintf(":%d", s.port)),
	)
}

// Get is a wrapper function that will add a new get handler to the server.
func (s *Server) Get(path string, response Response) {
	s.echo.GET(path, func(c echo.Context) error {
		return c.JSON(response.Code, response.Body)
	})
}

// Post is a wrapper function that will add a new get handler to the server.
func (s *Server) Post(path string, response Response) {
	s.echo.POST(path, func(c echo.Context) error {
		return c.JSON(response.Code, response.Body)
	})
}

// Put is a wrapper function that will add a new get handler to the server.
func (s *Server) Put(path string, response Response) {
	s.echo.PUT(path, func(c echo.Context) error {
		return c.JSON(response.Code, response.Body)
	})
}

// Patch is a wrapper function that will add a new get handler to the server.
func (s *Server) Patch(path string, response Response) {
	s.echo.PATCH(path, func(c echo.Context) error {
		return c.JSON(response.Code, response.Body)
	})
}

// Delete is a wrapper function that will add a new get handler to the server.
func (s *Server) Delete(path string, response Response) {
	s.echo.DELETE(path, func(c echo.Context) error {
		return c.JSON(response.Code, response.Body)
	})
}

// ReadJsonFile is a json wrapper that will read a file and convert the content
// to an interface{}. On fail it will panic.
func ReadJsonFile(path string) interface{} {
	// read file
	jsonFile, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("could not open %s, error: %s", path, err))
	}
	defer jsonFile.Close()
	fileContent, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(fmt.Sprintf("error reading contents in %s, error: %s", path, err))
	}

	// parse file to json
	var result interface{}
	err = json.Unmarshal(fileContent, &result)
	if err != nil {
		panic(fmt.Sprintf("could not convert file to json %s, error: %s", path, err))
	}
	return result
}
