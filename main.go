package main

func main() {
	mock := New()

	mock.Get("/", 200, struct {
		Number int
		Text   string
	}{
		42,
		"Hello world!",
	})

	mock.Get("/get", 200, map[string]interface{}{
		"hello": "test from a",
		"world": 42,
	})

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
