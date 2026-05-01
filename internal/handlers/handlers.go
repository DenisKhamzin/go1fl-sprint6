package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func SimpleGetHandler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
	fmt.Println("Запрос принят")
}

func Uploader(res http.ResponseWriter, req *http.Request) {
	err := req.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(res, "Error parsing form: "+err.Error(), http.StatusBadRequest)
		return
	}
	// 2. Получаем файл из формы с ключом "file"
	file, _, err := req.FormFile("myFile")
	if err != nil {
		http.Error(res, "Error retrieving file", http.StatusBadRequest)
		fmt.Println(req.Header)
		return
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(res, "Error reading file", http.StatusInternalServerError)
		fmt.Println("2")
		return
	}
	fmt.Println("Файл загружен")
	// Отправляем содержимое в ответ
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	res.Write(content)
}
