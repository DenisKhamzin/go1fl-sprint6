package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
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
		return
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(res, "Error reading file", http.StatusInternalServerError)
		return
	}
	fmt.Println("Файл загружен")
	// Отправляем содержимое в ответ
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	res.Write(content)
	// создаем файл
	text := service.Converter(string(content))
	err = fileCreator(text)
	if err != nil {
		fmt.Printf("Error: %s\n, err")
	}

}

func fileCreator(content string) error {
	now := time.Now()
	fileName := now.Format("2006-01-02 15:04:05") + ".txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Creating error: %s\n", err)
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("Writing error: %s\n", err)
	}
	return nil
}
