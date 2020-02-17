package action

import (
	"net/http"
	"service"
)

func UploadFile(writer http.ResponseWriter, r *http.Request) {
	err := service.Uploader(r)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(201)
	writer.Write([]byte("UPLOAD OK"))
}
