package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

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

// TestDatabase tests wether the database is initialized correctly.
func TestDatabase(t *testing.T) {
	// Initialize the test container database
	ctx := context.Background()
	container, database, err := InitTestDatabase(ctx)
	defer container.Terminate(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Check whether the database behaves correctly
	if "postgres" != database.Name() {
		t.Fatal("A problem occured while initializing the database")
	}

	if tx := database.Exec("select * from unknown"); tx.Error == nil {
		t.Fatal("The unknown table should not have been initialized")
	}

	if tx := database.Exec("select * from accounts"); tx.Error != nil {
		t.Fatal("The accounts table should have been initialized")
	}

	if tx := database.Exec("select * from tokens"); tx.Error != nil {
		t.Fatal("The tokens table should have been initialized")
	}

	if tx := database.Raw("select * from resources"); tx.Error != nil {
		t.Fatal("The resources table should have been initialized")
	}
}
