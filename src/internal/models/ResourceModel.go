package models

import "time"

type Resource struct {
	Uuid    string `gorm:"primary_key"`
	Document string
	AccountId string `gorm:"references:GoTrueId"`
	// Account Account `gorm:"references:GoTrueId"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

