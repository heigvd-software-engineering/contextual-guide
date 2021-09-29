package internal

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

	errorList := resource.ValidateResource()

	return resource,errorList
}
