package models

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"
)

const table string = "upload"

type Upload struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	OriginalName string `json:"original_name"`
	Extension    string `json:"extension"`
	Path         string `json:"path"`
	Filesize     int64  `json:"filesize"`
	created      time.Time
}

func CreateUpload(uploadData Upload) (id int64, err error) {

	var sqlStatement = fmt.Sprintf("INSERT INTO %s (name, original_name, extension, path, filesize) VALUES (?, ?, ?, ?, ?)", table)
	var valueArgs []interface{}
	id = 0

	valueArgs = append(valueArgs,
		uploadData.Name,
		uploadData.OriginalName,
		uploadData.Extension,
		uploadData.Path,
		strconv.FormatInt(uploadData.Filesize, 10))

	res, err := db.Exec(sqlStatement, valueArgs...)

	if err != nil {
		log.Println("Error inserting data in database: " + err.Error())
		return
	} else {
		id, err := res.LastInsertId()

		if err != nil {
			log.Println("Error when get last insert id: " + err.Error())
		} else {
			return id, nil
		}
	}

	return
}

func GetFile(id int64) *Upload {
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
