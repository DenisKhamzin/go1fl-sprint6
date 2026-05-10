package server

import (
	"fmt"
	"net/http"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

func HttpServer() error {
	// console success message that server is started
	fmt.Println("Server is starting")
	// handler for GET request
	http.HandleFunc("/", handlers.SimpleGetHandler)
	// handler for POST request
	http.HandleFunc("/upload", func(res http.ResponseWriter, req *http.Request) {
		// check if method is POST
		if req.Method != http.MethodPost {
			http.Error(res, "Wrong http-method", http.StatusMethodNotAllowed)
			return
		}
		handlers.Uploader(res, req)
	})
	// listening 8080 port with default router
	err := http.ListenAndServe(":8080", nil)
	return err
}
