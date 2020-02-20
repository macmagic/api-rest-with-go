package action

import (
	"apicommon"
	"net/http"
	"service"
)

func UploadFile(writer http.ResponseWriter, r *http.Request) {
	upload, err := service.Uploader(r)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	apicommon.JsonResponse(writer, map[string]interface{}{"content": upload}, http.StatusCreated)
}
