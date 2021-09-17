package repository

import (
	"main/src/internal/database"
	"main/src/internal/models"
)

type resourceRepository struct {
}

type IResourceRepository interface {
	GetResource(string) *models.Resource
	GetAllResource() []models.Resource
	CreateResource(*models.Resource) *models.Resource
}

var (
	ResourceRepository IResourceRepository
)

func init() {
	ResourceRepository = &resourceRepository{}
}

func (ur *resourceRepository) GetResource(id string) *models.Resource {
	var resource models.Resource
	database.DB.Where(&models.Resource{Id: id}).Find(&resource)
	return &resource
}

func (ur *resourceRepository) CreateResource(model *models.Resource) *models.Resource {
	database.DB.Create(model)
	return model
}

func (ur *resourceRepository) GetAllResource() []models.Resource {
	resources := []models.Resource{}
	database.DB.Find(&resources)
	return resources
}
