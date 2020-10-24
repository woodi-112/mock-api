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

	mock.Get("/a", 200, map[string]interface{}{
		"hello": "test from a",
		"world": 42,
	})

	mock.Start()
}
