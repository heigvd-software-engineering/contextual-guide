package repository

import (
	"fmt"
	"main/src/internal/database"
	"main/src/internal/models"
)

type resourceRepository struct {
}

type IResourceRepository interface {
	GetResource(string) *models.Resource
	GetAllResource() []models.Resource
	GetAllResourceByAccountId(string) []models.Resource
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
	database.DB.Where(&models.Resource{Uuid: id}).Find(&resource)
	return &resource
}

func (ur *resourceRepository) CreateResource(model *models.Resource) *models.Resource {
	database.DB.Create(model)
	return model
}

func (ur *resourceRepository) GetAllResource() []models.Resource {
	var resources []models.Resource
	database.DB.Preloads("Account").Find(&resources)
	fmt.Println(resources)
	return resources
}

func (ur *resourceRepository) GetAllResourceByAccountId(id string) []models.Resource {
	var resources []models.Resource
	database.DB.Preloads("Account").Find(&resources)
	database.DB.Where(&models.Resource{AccountId: id}).Find(&resources)
	fmt.Println(resources)
	return resources
}
