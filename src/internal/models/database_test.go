package models

import (
	"context"
	"testing"
)

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
