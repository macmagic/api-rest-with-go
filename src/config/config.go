package config

import (
	"encoding/json"
	"log"
	"os"
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

var Config *AppConfig
var once sync.Once

func GetAppConfig() *AppConfig {
	once.Do(func() {
		appServerConfig := loadConfig()
		Config = &appServerConfig
	})
	return Config
}

func loadConfig() AppConfig {
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

	return appConfig
}
