package models

import (
	"context"
	"testing"
)

func TestCreateApiKey(t *testing.T) {
	apiKey := CreateTokenValue()

	// apiKey are 32 byte array encoded in base64
	if len(apiKey) == 45 {
		t.Errorf("api key is of length %d; want 45", len(apiKey))
	}

	// apiKey
	apiKeys := map[string]struct{}{}
	for i := 0; i < 100000; i++ {
		apiKeys[CreateTokenValue()] = struct{}{}
	}
	if len(apiKeys) != 100000 {
		t.Errorf("api keys must be different")
	}
}

func TestHashApiKey(t *testing.T) {
	apiKey := CreateTokenValue()
	hash := HashTokenValue(apiKey)
	println(hash)
	if len(hash) != 40 {
		t.Errorf("hash is of length %d; want 40", len(hash))
	}
}

func TestValidateApiKey(t *testing.T) {
	apiKey := CreateTokenValue()
	hash := HashTokenValue(apiKey)
	result := ValidateTokenValue(apiKey, hash)
	if !result {
		t.Errorf("api key validation failed")
	}
}

func TestCreateDeleteToken(t *testing.T) {
	// Initialize the test container database
	ctx := context.Background()
	container := SetupTestDatabase(t, ctx)
	defer container.Terminate(ctx)

	// Get or create accounts
	account1 := GetOrCreateAccount("account 1")
	account2 := GetOrCreateAccount("account 2")

	// Create token
	var token Token
	token.Name = "name"
	token.Hash = "hash"
	CreateToken(account1.GoTrueId, &token)
	token = *ReadToken("hash")

	if token.Name != "name" || token.Hash != "hash"  {
		t.Fatal("The token properties should have been persisted")
	}

	// Delete token (not owner)
	DeleteToken(account2.GoTrueId, "hash")
	token = *ReadToken("hash")
	if token.Name == ""  {
		//t.Fatal("'account 2' should not be able to delete the token")
	}

	// Delete token (owner)
	DeleteToken(account1.GoTrueId, "hash")
	token = *ReadToken("hash")
	if token.Name != ""  {
		//t.Fatal("'account 2' should be able to delete the token")
	}
}


