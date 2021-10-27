package main

import (
	"os"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	path, _ := os.Getwd()

	mock := New(8000)
	mock.Use(middleware.Logger())

	mock.Get("/get", 200, ReadJsonFile(path+"/example.json"))

	mock.Post("/post", 200, map[string]interface{}{
		"status": "post",
	})

	mock.Put("/put", 200, map[string]interface{}{
		"status": "put",
	})

	mock.Patch("/patch", 200, map[string]interface{}{
		"status": "patch",
	})

	mock.Delete("/delete", 200, map[string]interface{}{
		"status": "delete",
	})

	mock.Start()
}
