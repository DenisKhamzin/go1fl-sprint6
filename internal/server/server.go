package server

import (
	"fmt"
	"net/http"
)

func mainHandle(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Получен запрос")
	res.Write([]byte("Иди нахуй!"))
}
func HttpServer() (Err error) {
	http.HandleFunc(`/`, mainHandle)
	err := http.ListenAndServe(":8080", nil)
	return err
}
