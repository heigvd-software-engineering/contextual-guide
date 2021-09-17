package models

type Token struct {
	Id    int64 `gorm:"primary_key;autoIncrement"`
	AccountId string `gorm:"primary_key"`
	Name  string `gorm:"unique"`
	Value string
}
