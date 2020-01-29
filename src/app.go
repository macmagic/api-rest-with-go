package src

import (
	config "./config"
	"log"
	"net/http"
	"strconv"
)

func Run() {
	router := config.LoadRouter()
	serverConfig := config.GetAppConfig()
	log.Println("Paso con el router")
	log.Println(serverConfig.Domain + ":" + strconv.Itoa(serverConfig.Port))
	server := http.ListenAndServe(serverConfig.Domain+":"+strconv.Itoa(serverConfig.Port), router)
	log.Fatal(server)
}
