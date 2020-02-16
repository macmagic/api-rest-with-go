package models

import (
	"log"
	"time"
)

const table string = "uploadfiles"

type Upload struct {
	Id           int
	Name         string
	OriginalName string
	Extension    string
	Path         string
	Filesize     int
	created      time.Time
}

func CreateUpload(uploadData Upload) {
	var sql = "INSERT INTO " + table + " (name, original_name, extension, path)"
	sql = sql + " VALUES ('" + uploadData.Name + "','" + uploadData.OriginalName + "', '" + uploadData.Extension + "','" + uploadData.Path + "')"

	insert, err := db.Query(sql)

	if err != nil {
		log.Fatal("Error inserting data in database: " + err.Error())
	}

	defer insert.Close()
}
