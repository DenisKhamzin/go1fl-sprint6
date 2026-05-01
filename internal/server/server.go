package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	//"github.com/go-chi/chi/v5/middleware"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

func HttpServer() (Err error) {
	router := chi.NewRouter()
	fmt.Println("Сервер стартанул")
	router.Get("/", handlers.SimpleGetHandler)
	router.Post("/upload", handlers.Uploader)

	err := http.ListenAndServe(":8080", router)
	return err
}
