package models

import (
	"config"
	"database/sql"
	"log"
	"strconv"
)
import _ "github.com/go-sql-driver/mysql"

var db *sql.DB

func init() {
	var appConfig = config.GetAppConfig()
	var dbHost = appConfig.Server.Database.Host
	var dbUser = appConfig.Server.Database.User
	var dbPassword = appConfig.Server.Database.Password
	var dbPort = appConfig.Server.Database.Port
	var dbName = appConfig.Server.Database.Db

	conn, err := sql.Open("mysql", dbUser+":"+dbPassword+"@tcp("+dbHost+":"+strconv.Itoa(dbPort)+")/"+dbName)

	if err != nil {
		log.Fatal("Error when try to connect to database: " + err.Error())
	}
	db = conn
	log.Println("Connection success")
}
