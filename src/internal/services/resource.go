package services

import (
	"main/src/internal/database"
	"main/src/internal/models"
)

func GetResource(id string) *models.Resource {
	var resource models.Resource
	database.DB.Where(&models.Resource{Uuid: id}).Find(&resource)
	return &resource
}

func CreateResource(model *models.Resource) *models.Resource {
	database.DB.Create(model)
	return model
}

func GetAllResource() []models.Resource {
	var resources []models.Resource
	database.DB.Preload("Account", &resources)
	return resources
}

func GetAllResourceByAccountId(id string) []models.Resource {
	var resources []models.Resource
	database.DB.Preload("Account", &resources)
	database.DB.Where(&models.Resource{AccountId: id}).Find(&resources)
	return resources
}
