package models

import (
	"time"
)

// Resource holds resource data in the database.
type Resource struct {
	Uuid string `gorm:"primary_key"`

	// The Attributes of the resource.
	Title       string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	Longitude   float32 `gorm:"not null"`
	Latitude    float32 `gorm:"not null"`
	Timestamp   time.Time
	Redirect    string

	// When an account is deleted, we must keep the associated resources.
	AccountId string
	Account   Account `gorm:"references:GoTrueId"`

	CreatedAt time.Time
	UpdatedAt time.Time
	// Resources shall never be deleted.
}

// ValidateResource return true if the attributes of the resource are valid.
func (r *Resource) ValidateResource() *ValidationError {
	errorList := make(ValidationError)

	notEmpty("title", r.Title, &errorList)
	notEmpty("description", r.Description, &errorList)

	inLatitudeBoundary("latitude", r.Latitude, &errorList)
	inLongitudeBoundary("longitude", r.Longitude, &errorList)

	isUrlFormat("redirect", r.Redirect, &errorList)

	return &errorList
}

func GetResource(id string) *Resource {
	var resource Resource
	DB.Where(&Resource{Uuid: id}).Find(&resource)
	return &resource
}

func CreateResource(model *Resource) *Resource {
	DB.Create(model)
	return model
}

func GetAllResource() []Resource {
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
