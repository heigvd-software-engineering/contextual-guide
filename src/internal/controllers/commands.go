package controllers

import (
	"github.com/lithammer/shortuuid/v3"
	"main/src/internal/models"
	"time"
)

type ResourceSaveCommand struct {
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	Timestamp        time.Time `json:"timestamp"`
	Longitude        float32   `json:"longitude"`
	Latitude         float32   `json:"latitude"`
	Redirect         string    `json:"redirect"`
	CustomProperties string    `json:"CustomProperties"`
}

func NewResource(command ResourceSaveCommand, accountId string) (*models.Resource, *models.ValidationError) {
	resource := &models.Resource{
		Uuid:        shortuuid.New(),
		Title:       command.Title,
		Description: command.Description,
		Timestamp:   command.Timestamp,
		Longitude:   command.Longitude,
		Latitude:    command.Latitude,
		Redirect:    command.Redirect,
		AccountId:   accountId,
	}

	errorList := resource.ValidateResource()

	return resource, errorList
}
