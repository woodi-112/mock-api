package main

import (
	"os"
)

func main() {
	path, _ := os.Getwd()

	mock := New(8000)

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
