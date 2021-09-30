package models

import "time"

// Account holds user data comming from GoTrue in the database.
type Account struct {
	GoTrueId  string     `gorm:"primaryKey"`
	Tokens    []Token    `gorm:"foreignKey:Account"`
	Resource  []Resource `gorm:"foreignKey:Resource"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}