package main

import (
	"fmt"
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	fmt.Println("Включаю сервер")
	err := server.HttpServer()
	if err != nil {
		log.Fatal(err)
	}
}
