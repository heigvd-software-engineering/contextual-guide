package internal

import (
	"testing"
)

func TestCreateApiKey(t *testing.T) {
	apiKey := CreateApiKey()

	// apiKey are 32 byte array encoded in base64
	if len(apiKey) == 45 {
		t.Errorf("api key is of length %d; want 45", len(apiKey))
	}

	// apiKey
	apiKeys := map[string]struct{}{}
	for i := 0; i < 100000; i++ {
		apiKeys[CreateApiKey()] = struct{}{}
	}
	if len(apiKeys) != 100000 {
		t.Errorf("api keys must be different")
	}
}

func TestHashApiKey(t *testing.T) {
	apiKey := CreateApiKey()
	hash := HashApiKey(apiKey)
	println(hash)
	if len(hash) != 40 {
		t.Errorf("hash is of length %d; want 40", len(hash))
	}
}

func TestValidateApiKey(t *testing.T) {
	apiKey := CreateApiKey()
	hash := HashApiKey(apiKey)
	result := ValidateApiKey(apiKey, hash)
	if !result {
		t.Errorf("api key validation failed")
	}
}


