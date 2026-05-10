package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// starting http server from server package
	err := server.HttpServer()
	if err != nil {
		// log possible error
		log.Fatal(err)
	}
}
