package models


type Token struct {
	Name  string `gorm:"unique"`
	Value string

	AccountId string
	Account Account `gorm:"references:GoTrueId"`
}
