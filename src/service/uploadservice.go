package service

import (
	"../config"
	"fmt"

)

var filePath string

/*func Uploader(r *http.Request) int {
	filePath = config.Config.Server.Files.Path
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("file")

	if err != nil {
		fmt.Println("Error retriving he file")
		fmt.Println(err)
		return 1
	}

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	defer file.Close()

	tmpFile, err := ioutil.TempFile(filePath, "upload-*.png")

	if err != nil {
		fmt.Println(err)
	}

	defer tmpFile.Close()

	fileBytes, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println(err)
	}

	tmpFile.Write(fileBytes)
	fmt.Println("Finish")

	return 0
}*/

func NewService(){
	fmt.Println(BaseConfig.Server.Domain)
}
