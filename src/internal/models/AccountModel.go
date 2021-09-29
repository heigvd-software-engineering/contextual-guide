package models

import (
	"time"
)

type Account struct {
	GoTrueId string `gorm:"primaryKey"`
	Tokens []Token `gorm:"foreignKey:Account"`
	Resource []Resource `gorm:"foreignKey:Resource"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
