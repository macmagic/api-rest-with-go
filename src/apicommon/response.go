package apicommon

import (
	"encoding/json"
	"net/http"
)

type ResponseData struct {
	Data interface{} `json:"content"`
	Meta interface{} `json:"meta"`
}

func JsonResponse(w http.ResponseWriter, data map[string]interface{}, httpStatus int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
