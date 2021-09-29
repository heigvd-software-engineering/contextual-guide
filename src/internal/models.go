package internal

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"time"
)

// Account holds user data comming from GoTrue in the database.
type Account struct {
	GoTrueId  string     `gorm:"primaryKey"`
	Tokens    []Token    `gorm:"foreignKey:Account"`
	Resource  []Resource `gorm:"foreignKey:Resource"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// Token holds hashed api key data  in the database.
type Token struct {
	Hash string `gorm:"primaryKey"`
	Name string

	// When an account is deleted, we must delete all the associated tokens
	AccountId string
	Account   Account `gorm:"references:GoTrueId,constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// CreateApiKey returns an random api key made of 32 bytes encoded in base64.
// The api keys are created with a cryptographically secure random number generator.
func CreateApiKey() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	string := base64.StdEncoding.EncodeToString(bytes)
	return string
}

// HashApiKey hashes a randomly generated api key with sha1 and returns an hexadecimal string.
// As the entropy of the randomly generated api key is high, we don't need to add salt.
func HashApiKey(apiKey string) string {
	hash := sha1.New()
	hash.Write([]byte(apiKey))
	sha := hash.Sum(nil)
	return hex.EncodeToString(sha)
}

// ValidateApiKey verifies that the provided api key corresponds to the provided hash.
func ValidateApiKey(apiKey string, hash string) bool {
	return HashApiKey(apiKey) == hash
}

// Resource holds resource data in the database.
type Resource struct {
	Uuid string `gorm:"primary_key"`

	// The Attributes of the resource.
	Title       string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	Longitude   float32 `gorm:"not null"`
	Latitude    float32 `gorm:"not null"`
	Timestamp   time.Time
	Redirect    string

	// When an account is deleted, we must keep the associated resources.
	AccountId string
	Account   Account `gorm:"references:GoTrueId"`

	CreatedAt time.Time
	UpdatedAt time.Time
	// Resources shall never be deleted.
}

// ValidateResource return true if the attributes of the resource are valid.
func (r *Resource) ValidateResource() *ValidationError {
	errorList := make(ValidationError)

	notEmpty("title", r.Title, &errorList)
	notEmpty("description", r.Description, &errorList)

	inLatitudeBoundary("latitude", r.Latitude, &errorList)
	inLongitudeBoundary("longitude", r.Longitude, &errorList)

	isUrlFormat("redirect", r.Redirect, &errorList)

	return &errorList
}
