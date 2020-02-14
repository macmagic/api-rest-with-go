package action

import (
	"net/http"
	"service"
)

func UploadFile(writer http.ResponseWriter, r *http.Request) {
	service.Uploader(r)
	writer.WriteHeader(201)
	writer.Write([]byte("UPLOAD OK"))
}
