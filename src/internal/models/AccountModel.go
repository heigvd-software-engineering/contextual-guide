package models

import "time"

type Account struct {
	GoTrueId string `gorm:"primaryKey"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
