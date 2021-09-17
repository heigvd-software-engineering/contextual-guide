package models

type Resource struct {
	Id    string `gorm:"primary_key"`
	Content string
}
