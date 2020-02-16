package service

import (
	"fmt"
	"io"
	"log"
	"models"
	"net/http"
	"os"
	"path/filepath"
)

var filePath string

func Uploader(r *http.Request) int {
	filePath = BaseConfig.Server.Files.Path
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("file")

	if err != nil {
		fmt.Println("Error retriving the file")
		fmt.Println(err)
		log.Fatal("Error when retreiving the file :(")
		return 1
	}

	path, _ := os.Getwd()

	path = path + "/../" + filePath

	log.Println(path)

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	defer file.Close()
	log.Println(path + handler.Filename)
	f, err := os.OpenFile(path+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		log.Fatal(err)
	}

	var uploadData = models.Upload{
		Id:           0,
		Name:         handler.Filename,
		OriginalName: handler.Filename,
		Extension:    filepath.Ext(handler.Filename),
		Path:         path + handler.Filename,
	}

	models.CreateUpload(uploadData)

	defer f.Close()

	io.Copy(f, file)

	return 0
}

func saveUpload(filename string, originalname string, extension string, path string) {

}
