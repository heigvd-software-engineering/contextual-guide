package models

import (
	"github.com/lithammer/shortuuid/v3"
	"gorm.io/gorm"
	"time"
)

// Resource holds resource data in the database.
type Resource struct {
	Uuid string `gorm:"primary_key" json:"uuid,omitempty"`

	// The Attributes of the resource.
	Title       string  `form:"title" json:"title" binding:"required"`
	Description string  `form:"description" json:"description,omitempty"`
	Longitude   float32 `form:"longitude" json:"longitude" binding:"required,gte=-180,lte=180"`
	Latitude    float32 `form:"latitude" json:"latitude" binding:"required,gte=-90,lte=90"`
	Timestamp   time.Time    `form:"timestamp" json:"timestamp,omitempty"`
	Redirect    string  `form:"redirect" json:"redirect,omitempty"`

	// When an account is deleted, we must keep the associated resources.
	AccountId string
	Account   Account `gorm:"references:GoTrueId"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func ReadResource(uuid string) *Resource {
	var resource Resource
	DB.Where(&Resource{Uuid: uuid}).Find(&resource)
	return &resource
}

func CreateResource(accountId string, resource *Resource) *Resource {
	resource.Uuid = shortuuid.New()
	resource.AccountId = accountId
	DB.Create(resource)
	return resource
}

func UpdateResource(accountId string, resource *Resource) *Resource {
	resource.AccountId = accountId
	DB.Where("account_id = ? and uuid = ?", accountId, resource.Uuid).Updates(resource)
	return resource
}

func DeleteResource(accountId string, uuid string) {
	DB.Where("account_id = ? and uuid = ?", accountId, uuid).Delete(&Resource{})
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
