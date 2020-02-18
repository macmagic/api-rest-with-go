package models

import (
	"database/sql"
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

	var sqlStatement = fmt.Sprintf("INSERT INTO %s (name, original_name, extension, path, filesize) VALUES (?, ?, ?, ?, ?)", table)
	var valueArgs []interface{}

	valueArgs = append(valueArgs,
		uploadData.Name,
		uploadData.OriginalName,
		uploadData.Extension,
		uploadData.Path,
		strconv.FormatInt(uploadData.Filesize, 10))

	_, err := db.Exec(sqlStatement, valueArgs...)

	if err != nil {
		log.Println("Error inserting data in database: " + err.Error())
		return err
	}

	return nil
}

func GetFile(id int) *Upload {
	var sqlStatement = fmt.Sprintf("SELECT id, name, original_name, extension, path, filesize FROM %s WHERE id=?;", table)
	var upload Upload
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&upload.Id, &upload.Name, &upload.OriginalName, &upload.Extension, &upload.Path, &upload.Filesize)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Not found!")
		return nil
	case nil:
		return &upload
	default:
		log.Panicln("Error! " + err.Error())
	}

	return nil
}

func GetFiles() []Upload {
	var sqlStatement = fmt.Sprintf("SELECT id, name, original_name, extension, path, filesize FROM %s;", table)
	var uploadRow Upload
	var uploads []Upload
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Panicln("Error when execute query in database: " + err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&uploadRow.Id, &uploadRow.Name, &uploadRow.OriginalName, &uploadRow.Extension, &uploadRow.Path, &uploadRow.Filesize)
		if err != nil {
			log.Panicln("Error when retrieve data from resultset: " + err.Error())
		}

		uploads = append(uploads, uploadRow)
	}

	return uploads
}
