package action

import (
	service "../service"
	"net/http"
)

func UploadFile(writer http.ResponseWriter, r *http.Request) {
	service.Uploader(r)
}
