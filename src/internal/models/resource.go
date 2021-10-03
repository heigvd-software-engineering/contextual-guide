package models

import (
	"gorm.io/gorm"
	"time"
)

// Resource holds resource data in the database.
type Resource struct {
	Uuid string `gorm:"primary_key"`

	// The Attributes of the resource.
	Title       string    `form:"title" json:"title" binding:"required"`
	Description string    `form:"description" json:"description"`
	Longitude   float32   `form:"longitude" json:"longitude" binding:"required,gte=-180,lte=180"`
	Latitude    float32   `form:"latitude" json:"latitude" binding:"required,gte=-90,lte=90"`
	Timestamp   time.Time `form:"timestamp" json:"timestamp"`
	Redirect    string    `form:"redirect" json:"redirect"`

	// When an account is deleted, we must keep the associated resources.
	AccountId string
	Account   Account `gorm:"references:GoTrueId"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Validate return true if the attributes of the resource are valid.
func (r *Resource) Validate() *ValidationError {
	errorList := make(ValidationError)

	notEmpty("title", r.Title, &errorList)
	notEmpty("description", r.Description, &errorList)

	inLatitudeBoundary("latitude", r.Latitude, &errorList)
	inLongitudeBoundary("longitude", r.Longitude, &errorList)

	isUrlFormat("redirect", r.Redirect, &errorList)

	return &errorList
}

func CreateResource(resource *Resource) *Resource {
	DB.Create(resource)
	return resource
}

func ReadResource(uuid string) *Resource {
	var resource Resource
	DB.Where(&Resource{Uuid: uuid}).Find(&resource)
	return &resource
}

func UpdateResource(model *Resource) *Resource {
	DB.Updates(model)
	return model
}

func DeleteResource(uuid string) {
	DB.Where("uuid = ?", uuid).Delete(&Resource{})
}

func GetAllResources() []Resource {
	var resources []Resource
	DB.Preload("Account", &resources)
	return resources
}

func GetAllResourceByAccountId(id string) []Resource {
	var resources []Resource
	DB.Preload("Account", &resources)
	DB.Where(&Resource{AccountId: id}).Find(&resources)
	return resources
}
