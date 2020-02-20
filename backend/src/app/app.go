package app

import (
	"config"
	"fmt"
	"log"
	"net/http"
	"router"
	"service"
	"strconv"
)

func Run() {
	log.Println("Load app config...")
	config.GetAppConfig()
	service.GetConfig(config.Config)
	log.Println("App config OK")

	log.Println("Generate routes")
	routerConfig := router.LoadRoutes()
	log.Println("Generate routes OK")

	log.Println(fmt.Sprintf("Load HTTP server in http://%s:%s", config.Config.Server.Domain, strconv.Itoa(config.Config.Server.Port)))
	server := http.ListenAndServe(config.Config.Server.Domain+":"+strconv.Itoa(config.Config.Server.Port), routerConfig)
	log.Fatal(server)
}
