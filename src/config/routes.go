package config

import (
	"action"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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

var functions = map[string]http.HandlerFunc{
	"Index":      action.IndexAction,
	"UploadFile": action.UploadFile,
}

func LoadRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	routes := getRoutesFromFile()
	for _, route := range routes.Routes {
		log.Print(route.Name)
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(functions[route.HandleFunc])
	}
	return router
}

func getRoutesFromFile() Routes {
	var routes Routes
	path, _ := os.Getwd()
	log.Println("Path of file: " + path + "/../config/routes.json")
	file, errOs := os.OpenFile(path+"/../config/routes.json", os.O_RDONLY, 0777)

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
