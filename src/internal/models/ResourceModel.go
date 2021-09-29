package models

import (
	"github.com/lithammer/shortuuid/v3"
	"time"
)


type ResourceSaveCommand struct {
	Title    string `json:"title"`
	Description    string `json:"description"`
	Timestamp    time.Time `json:"timestamp"`
	Longitude    float32 `json:"longitude"`
	Latitude    float32 `json:"latitude"`
	Redirect    string `json:"redirect"`
	CustomProperties string `json:"CustomProperties"`
}

type Resource struct {
	Uuid    string `gorm:"primary_key"`

	// The Attributes
	Title string `gorm:"not null"`
	Description string `gorm:"not null"`
	Longitude float32 `gorm:"not null"`
	Latitude float32 `gorm:"not null"`
	Timestamp time.Time
	Redirect string

	// When an account is deleted, we must keep the associated resources
	AccountId string
	Account Account `gorm:"references:GoTrueId"`

	CreatedAt time.Time
	UpdatedAt time.Time
	// Resources shall never be deleted
}


func NewResource(command ResourceSaveCommand, accountId string) (*Resource, *ValidationError){
	resource := &Resource{
		Uuid: shortuuid.New(),
		Title: command.Title,
		Description: command.Description,
		Timestamp: command.Timestamp,
		Longitude: command.Longitude,
		Latitude: command.Latitude,
		Redirect:  command.Redirect,
		AccountId: accountId,
	}

	errorList := resource.Validate()

	return resource,errorList
}

func (r *Resource) Validate() *ValidationError {

	errorList := make(ValidationError)

	notEmpty("title",r.Title,&errorList)
	notEmpty("description",r.Description,&errorList)

	inLatitudeBoundary("latitude", r.Latitude, &errorList)
	inLongitudeBoundary("longitude", r.Longitude, &errorList)

	isUrlFormat("redirect",r.Redirect,&errorList)

	return &errorList
}
