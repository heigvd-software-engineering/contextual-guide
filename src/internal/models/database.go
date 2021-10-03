package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

// DB is a global variable for the gorm database.
var DB *gorm.DB

// ConnectDatabaseEnv connects to the database using environment variables.
func ConnectDatabaseEnv() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"))
	ConnectDatabase(dsn)
}

// ConnectDatabase connects to the database using a data source name
func ConnectDatabase(dsn string) {
	db, err := InitDatabase(dsn)
	if err != nil {
		panic(err)
	}
	DB = db
}

// InitDatabase opens and migrates the database using a data source name
func InitDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Migrate the db
	if err == nil {
		db.AutoMigrate(&Account{})
		db.AutoMigrate(&Resource{})
		db.AutoMigrate(&Token{})
	}

	return db, err
}