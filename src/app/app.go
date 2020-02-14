package app

import (
	"config"
	"log"
	"net/http"
	"service"
	"strconv"
)

func Run() {
	log.Println("Generate router")
	config.GetAppConfig()
	service.GetConfig(config.Config)
	router := config.LoadRouter()
	log.Println("Get app config")
	log.Println("Paso con el router")
	log.Println(config.Config.Server.Domain + ":" + strconv.Itoa(config.Config.Server.Port))
	server := http.ListenAndServe(config.Config.Server.Domain+":"+strconv.Itoa(config.Config.Server.Port), router)
	log.Fatal(server)
}
