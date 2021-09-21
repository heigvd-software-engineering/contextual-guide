package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"main/src/internal/models"
	"os"
)

var (
	DB *gorm.DB
)

func init() {
	DB, _ = connect()
}

func connect() (*gorm.DB, error) {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open("postgres", connString)

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Account{})
	db.AutoMigrate(&models.Resource{})
	db.AutoMigrate(&models.Token{})

	return db, nil
}
