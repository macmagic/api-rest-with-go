package service

import "models"

func GetFile(id int) *models.Upload {
	return models.GetFile(int64(id))
}

func GetFiles() []models.Upload {
	return models.GetFiles()
}
