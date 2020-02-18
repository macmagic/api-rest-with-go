package action

import (
	"apicommon"
	"net/http"
	"service"
)

func GetFiles(writer http.ResponseWriter, r *http.Request) {
	uploads := service.GetFiles()
	apicommon.JsonResponse(writer, map[string]interface{}{"content": uploads}, http.StatusOK)
}
