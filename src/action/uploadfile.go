package action

import (
	"apicommon"
	"net/http"
	"service"
)

func UploadFile(writer http.ResponseWriter, r *http.Request) {
	err := service.Uploader(r)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	apicommon.JsonResponse(writer, nil, http.StatusCreated)
}
