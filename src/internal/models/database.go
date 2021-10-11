package models

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"testing"
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


// InitTestDatabase initializes a container and a database connection.
func InitTestDatabase(ctx context.Context) (testcontainers.Container, *gorm.DB, error) {
	// Create the Postgres test container
	req := testcontainers.ContainerRequest{
		Image:        "postgis/postgis:latest",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_DB": "postgres",
			"POSTGRES_USER": "postgres",
			"POSTGRES_PASSWORD": "postgres",
		},
		WaitingFor: wait.ForLog("database system is ready to accept connections").WithOccurrence(2),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		panic(err)
	}

	// Get the host
	host, err := container.Host(ctx)
	if err != nil {
		container.Terminate(ctx)
		return nil, nil, err
	}

	// Get the port
	port, err := container.MappedPort(ctx, "5432")
	if err != nil {
		container.Terminate(ctx)
		return nil, nil, err
	}

	// Create connection string to the test container
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port.Port(),
		"postgres",
		"postgres",
		"postgres")

	// Connect to the database
	database, err := InitDatabase(dsn)
	if err != nil {
		container.Terminate(ctx)
		return nil, nil, err
	}

	return container, database, nil
}

// SetupTestDatabase setups the global DB variable.
func SetupTestDatabase(t *testing.T, ctx context.Context) testcontainers.Container {
	container, database, err := InitTestDatabase(ctx)
	if err != nil {
		t.Fatal(err)
	}
	DB = database
	return container
}