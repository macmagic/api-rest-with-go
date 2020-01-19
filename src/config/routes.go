package config

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"os"
)

type Route struct {
	Name       string
	Method     string
	Path       string
	HandleFunc string
}

type Routes struct {
	Routes []Route
}

func LoadRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	/*routes := getRoutesFromFile()
	for()*/
	return router
}

func getRoutesFromFile() Routes {
	var routes Routes
	path, _ := os.Getwd()
	fmt.Println(path)
	file, errOs := os.OpenFile(path+"/config/routes.json", os.O_RDONLY, 0777)

	if errOs != nil {
		log.Fatal(errOs)
	}

	decoder := json.NewDecoder(file)

	err := decoder.Decode(&routes)

	if err != nil {
		log.Fatal(err)
	}
	return routes
}
