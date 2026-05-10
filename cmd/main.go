package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// creating new logger
	logger := log.New(os.Stdout, "http-server", log.LstdFlags|log.Lshortfile)
	// creating new http server using current logger
	Server := server.NewServer(logger)

	// server starts listening port
	err := Server.HttpServer.ListenAndServe()
	// logging error
	if err != nil {
		logger.Fatal("HTTP server doesn't start: ", err)
	}
}
