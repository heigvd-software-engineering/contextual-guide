package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open("mysql", connString)

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Account{})
	db.AutoMigrate(&models.Resource{})
	db.AutoMigrate(&models.Token{})

	return db, nil
}
