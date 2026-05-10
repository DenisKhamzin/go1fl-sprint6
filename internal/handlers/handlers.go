package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// handler for GET requests
func SimpleGetHandler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
}

// handler for POST requests to /upload
func Uploader(res http.ResponseWriter, req *http.Request) {
	// parsing data from request body
	err := req.ParseMultipartForm(10 << 20)
	// error check
	if err != nil {
		http.Error(res, "Error parsing form: ", http.StatusBadRequest)
		return
	}
	// getting file from field "myFile"
	file, _, err := req.FormFile("myFile")
	//error check
	if err != nil {
		http.Error(res, "Error retrieving file: ", http.StatusBadRequest)
		return
	}
	// closing uploaded file
	defer file.Close()

	// reading uploaded file
	content, err := io.ReadAll(file)
	// error check
	if err != nil {
		http.Error(res, "Error reading file: ", http.StatusInternalServerError)
		return
	}

	// converting file content to morse or text depending on it
	text := service.Converter(string(content))

	// creating file and write converted data into it
	err = fileCreator(text)
	// error check
	if err != nil {
		http.Error(res, "Error creating file: ", http.StatusInternalServerError)
	}
	// console success message that file is created
	fmt.Println("file created")
	// preparing response
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	// replying converted text
	res.Write([]byte(text))
}

// dedicated function for creating local file
func fileCreator(content string) error {
	now := time.Now()
	fileName := now.Format("2006-01-02 15:04:05") + ".txt"
	// creating file
	file, err := os.Create(fileName)
	// error check
	if err != nil {
		fmt.Printf("Creating error: %s\n", err)
	}
	// closing created file
	defer file.Close()
	// trying to write converted text into file and check error
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("Writing error: %s\n", err)
	}
	return nil
}
