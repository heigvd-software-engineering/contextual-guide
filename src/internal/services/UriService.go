package services

import (
	"main/src/internal/models"
	"main/src/internal/repository"
)

type uriService struct {

}

var (
	UriService uriService
)

func (us *uriService) CreateUri(newUri *models.Uri) *models.Uri  {

	uri := repository.UriRepository.CreateUri(newUri)
	return uri
}

func (us *uriService) GetAll() []models.Uri  {

	uri := repository.UriRepository.GetAll()
	return uri
}

func (us *uriService) GetOne(uuid string) *models.Uri  {

	uri := repository.UriRepository.GetUri(uuid)
	return uri
}