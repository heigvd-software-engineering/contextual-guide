package repository

import (
	"main/src/internal/database"
	"main/src/internal/models"
)

type uriRepository struct {
}

type IUriRepository interface {
	GetUri(string) *models.Uri
	GetAll() []models.Uri
	CreateUri(*models.Uri) *models.Uri
}

var (
	UriRepository IUriRepository
)

func init() {
	UriRepository = &uriRepository{}
}

func (ur *uriRepository) GetUri(uuid string) *models.Uri {

	var uris models.Uri
	database.DB.Where(&models.Uri{Id: uuid}).Find(&uris)

	return &uris
}

func (ur *uriRepository) CreateUri(model *models.Uri) *models.Uri {
	database.DB.Create(model)

	return model
}

func (ur *uriRepository) GetAll() []models.Uri {

	uris := []models.Uri{}

	database.DB.Find(&uris)
	return uris
}
