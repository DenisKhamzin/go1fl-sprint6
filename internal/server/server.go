package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// http server struct
type Server struct {
	Logger     *log.Logger
	HttpServer *http.Server
}

// function for creating server instance
func NewServer(logger *log.Logger) *Server {
	// creating router
	router := http.NewServeMux()
	// register handlers in current router
	router.HandleFunc("/", handlers.SimpleGetHandler)
	router.HandleFunc("/upload", handlers.Uploader)

	// creating instance if server with requiremtnt params
	serverHttp := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	// returning created instance of http server
	return &Server{
		Logger:     logger,
		HttpServer: serverHttp,
	}
}
