package models

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

const table string = "uploadfiles"

type Upload struct {
	Id           int
	Name         string
	OriginalName string
	Extension    string
	Path         string
	Filesize     int64
	created      time.Time
}

func CreateUpload(uploadData Upload) error {

	var sql = fmt.Sprintf("INSERT INTO %s (name, original_name, extension, path, filesize) VALUES (?, ?, ?, ?, ?)", table)
	var valueArgs []interface{}

	valueArgs = append(valueArgs,
		uploadData.Name,
		uploadData.OriginalName,
		uploadData.Extension,
		uploadData.Path,
		strconv.FormatInt(uploadData.Filesize, 10))

	_, err := db.Exec(sql, valueArgs...)

	if err != nil {
		log.Println("Error inserting data in database: " + err.Error())
		return err
	}

	return nil
}
