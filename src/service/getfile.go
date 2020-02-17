package service

import "models"

func GetFile(id int) *models.Upload {
	return models.GetFile(id)
}
