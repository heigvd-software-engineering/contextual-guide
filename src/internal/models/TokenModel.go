package models

type Token struct {
	Id    int64 `gorm:"primary_key;autoIncrement"`
	Name  string
	Value string
}
