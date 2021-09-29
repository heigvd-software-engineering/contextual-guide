package models

import "time"

type Token struct {
	Hash string `gorm:"primaryKey"`
	Name  string

	// When an account is deleted, we must delete all the associated tokens
	AccountId string
	Account Account `gorm:"references:GoTrueId,constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

}
