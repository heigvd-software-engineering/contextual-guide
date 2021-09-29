package internal

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

// DB is a global variable for the gorm database.
var DB *gorm.DB

// ConnectDatabase connects to the database using environment variables.
func ConnectDatabase() {
	// Open the postgresql connection
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrate the db
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Resource{})
	db.AutoMigrate(&Token{})

	DB = db
}
