package models

import (
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


