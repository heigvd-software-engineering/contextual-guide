package models

type Resource struct {
	Uuid    string `gorm:"primary_key"`
	Content string

	AccountId string
	Account Account `gorm:"references:GoTrueId"`
}


