package models

import (
	"time"
)

// Account holds user data comming from GoTrue in the database.
type Account struct {
	GoTrueId  string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// GetOrCreateAccount gets or creates an account corresponding to the provided GoTrue identifier.
func GetOrCreateAccount(goTrueId string) *Account {
	var model Account
	DB.Where(&Account{GoTrueId: goTrueId}).FirstOrCreate(&model)
	return &model
}
