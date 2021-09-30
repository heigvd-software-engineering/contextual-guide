package models

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"time"
)

// Token holds hashed api key data  in the database.
type Token struct {
	Hash string `gorm:"primaryKey"`
	Name string

	// When an account is deleted, we must delete all the associated tokens
	AccountId string
	Account   Account `gorm:"references:GoTrueId"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func GetToken(id int64) *Token {
	var tokens Token
	//database.DB.Where(&Token{: id}).Find(&tokens)
	return &tokens
}

func CreateToken(model *Token) *Token {
	DB.Create(model)
	return model
}

// CreateTokenValue returns an random token value made of 32 bytes encoded in base64.
// The token values are created with a cryptographically secure random number generator.
func CreateTokenValue() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return base64.StdEncoding.EncodeToString(bytes)
}

// HashTokenValue hashes a randomly generated token value with sha1 and returns an hexadecimal string.
// As the entropy of the randomly generated token value is high, we don't need to add salt.
func HashTokenValue(value string) string {
	hash := sha1.New()
	hash.Write([]byte(value))
	sha := hash.Sum(nil)
	return hex.EncodeToString(sha)
}

// ValidateTokenValue verifies that the provided token value corresponds to the provided hash.
func ValidateTokenValue(value string, hash string) bool {
	return HashTokenValue(value) == hash
}

func ListTokenByAccountId(id string) []Token {
	tokens := []Token{}
	DB.Where(&Token{AccountId: id}).Find(&tokens)
	return tokens
}

func DeleteToken(id int64) {
	DB.Delete(&Token{}, id)
}

func GetTokenByValue(value string) *Token {
	var tokens Token
	DB.Where(&Token{Hash: value}).Find(&tokens)
	return &tokens
}



