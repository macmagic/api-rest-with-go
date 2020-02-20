package action

import (
	"apicommon"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"service"
	"strconv"
)

func GetFile(writer http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	paramId, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println("Error when try to convert string to int")
		apicommon.JsonResponse(writer, nil, http.StatusBadRequest)
		return
	}

	upload := service.GetFile(paramId)

	if upload == nil {
		apicommon.JsonResponse(writer, nil, http.StatusNotFound)
		return
	}

	apicommon.JsonResponse(writer, map[string]interface{}{"content": upload}, http.StatusOK)
}
