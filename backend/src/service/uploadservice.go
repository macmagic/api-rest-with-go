package service

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"models"
	"net/http"
	"os"
	"path/filepath"
)

const fileLength = 20

func Uploader(r *http.Request) (upload *models.Upload, err error) {
	uploadPath := BaseConfig.Server.Files.Path
	filePath := BaseConfig.Server.Files.ProjectRootPath + uploadPath

	generateAppDirectories(filePath)

	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("file")

	if err != nil {
		fmt.Println("Error retrieving the file")
		fmt.Println(err)
		log.Fatal("Error when retrieving the file :(")
		return nil, err
	}

	extension := filepath.Ext(handler.Filename)

	defer file.Close()

	systemFileName := generateRandomName(fileLength) + extension
	f, err := os.OpenFile(filePath+systemFileName, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer f.Close()

	io.Copy(f, file)

	fi, err := f.Stat()

	if err != nil {
		log.Fatal("Error when get stat from file: " + err.Error())
		return nil, err
	}

	id, err := saveUpload(systemFileName, handler.Filename, extension, uploadPath, fi.Size())

	if err != nil {
		return nil, err
	}

	uploadData := models.GetFile(id)

	return uploadData, nil
}

func saveUpload(filename string, originalname string, extension string, path string, filesize int64) (id int64, err error) {
	var uploadData = models.Upload{
		Id:           0,
		Name:         filename,
		OriginalName: originalname,
		Extension:    extension,
		Path:         path,
		Filesize:     filesize,
	}

	return models.CreateUpload(uploadData)
}

func generateRandomName(length int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz1234567890")

	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func generateAppDirectories(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.Mkdir(filePath, 0777)
	}
}
