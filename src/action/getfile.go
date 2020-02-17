package action

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"service"
	"strconv"
)

func GetFile(writer http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Panicln("Error when try to convert string to int")
	}

	upload := service.GetFile(id)

	if upload == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(upload)
}
