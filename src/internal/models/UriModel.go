package models

type Uri struct {
	Id       string `gorm:"primary_key"`
	Document string
}
