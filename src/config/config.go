package config

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"sync"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Db       string
}

type FilesConfig struct {
	Path string
}

type ServerConfig struct {
	Port     int
	Domain   string
	Database DatabaseConfig
	Files    FilesConfig
}

type AppConfig struct {
	Server ServerConfig
}

var appConfig *ServerConfig
var once sync.Once

func GetAppConfig() *ServerConfig {
	once.Do(func() {
		appServerConfig := loadConfig()
		appConfig = &appServerConfig
	})
	return appConfig
}

func loadConfig() ServerConfig {
	var appConfig AppConfig
	path, _ := os.Getwd()
	file, errOs := os.OpenFile(path+"/config/app.json", os.O_RDONLY, 0777)

	if errOs != nil {
		log.Fatal(errOs)
	}

	decoder := json.NewDecoder(file)

	err := decoder.Decode(&appConfig)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(appConfig.Server.Domain + ":" + strconv.Itoa(appConfig.Server.Port))
	return appConfig.Server
}
