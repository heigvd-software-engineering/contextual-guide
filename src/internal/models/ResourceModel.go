package models

import (
	"time"
)

type Resource struct {
	Uuid    string `gorm:"primary_key"`

	Title string
	Description string
	Timestamp time.Time
	Longitude float32
	Latitude float32
	Redirect string

	CustomProperties string

	AccountId string `gorm:"references:GoTrueId"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

