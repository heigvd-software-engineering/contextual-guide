package models

import (
	"context"
	"fmt"
	"testing"
)

func TestCreateUpdateDeleteResource(t *testing.T) {
	// Initialize the test container database
	ctx := context.Background()
	container := SetupTestDatabase(t, ctx)
	defer container.Terminate(ctx)

	// Get or create accounts
	account1 := GetOrCreateAccount("account 1")
	account2 := GetOrCreateAccount("account 2")

	// Create resource
	var resource Resource
	resource.Title = "title 1"
	CreateResource(account1.GoTrueId, &resource)
	if len(resource.Uuid) != 22  {
		t.Fatal("The resource's uuid should be of length 22")
	}
	if resource.AccountId != "account 1" {
		t.Fatal("The resource's account should be 'account 1'")
	}
	if resource.Title != "title 1" {
		t.Fatal("The resource's title should be 'title 1'")
	}

	// Update resource (not owner)
	resource.Title = "title 2"
	UpdateResource(account2.GoTrueId, &resource)
	resource = *ReadResource(resource.Uuid)
	if resource.Title == "title 2" {
		t.Fatal("'account 2' should not be able to update the resource")
	}

	// Update resource (owner)
	resource.Title = "title 2"
	UpdateResource(account1.GoTrueId, &resource)
	resource = *ReadResource(resource.Uuid)
	if resource.Title != "title 2" {
		t.Fatal("'account 1' should be able to update the resource")
	}

	// Delete resource (not owner)
	DeleteResource(account2.GoTrueId, resource.Uuid)
	resource = *ReadResource(resource.Uuid)
	if resource.Uuid == ""  {
		t.Fatal("'account 2' should not be able to delete the resource")
	}

	// Delete resource (owner)
	DeleteResource(account1.GoTrueId, resource.Uuid)
	fmt.Println(resource.DeletedAt)
	if resource.Uuid != ""  {
		t.Fatal("'account 2' should be able to delete the resource")
	}
}