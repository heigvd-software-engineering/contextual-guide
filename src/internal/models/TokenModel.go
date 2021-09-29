package models

import "time"

type Token struct {
	Value string `gorm:"primaryKey"`

	Name  string
	AccountId string
	Account Account `gorm:"references:GoTrueId"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
