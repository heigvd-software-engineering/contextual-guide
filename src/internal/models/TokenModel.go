package models

import "time"

type Token struct {
	Name  string `gorm:"unique"`
	Value string

	AccountId string
	Account Account `gorm:"references:GoTrueId"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
