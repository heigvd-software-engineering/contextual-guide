package services

import (
	"main/src/internal/models"
	"main/src/internal/repository"
)

type resourceService struct {
}

var (
	ResourceService resourceService
)

func (us *resourceService) CreateResource(newResource *models.Resource) *models.Resource {
	resource := repository.ResourceRepository.CreateResource(newResource)
	return resource
}

func (us *resourceService) GetAll() []models.Resource {
	resource := repository.ResourceRepository.GetAllResource()
	return resource
}

func (us *resourceService) GetOne(id string) *models.Resource {
	resource := repository.ResourceRepository.GetResource(id)
	return resource
}
