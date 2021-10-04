package models

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"gorm.io/gorm"
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
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func ListTokenByAccountId(accountId string) []Token {
	tokens := []Token{}
	DB.Where("account_id = ?", accountId).Find(&tokens)
	return tokens
}

func ReadToken(hash string) *Token {
	var token Token
	DB.Where("hash = ?", hash).Find(&token)
	return &token
}

func CreateToken(accountId string, token *Token) *Token {
	token.AccountId = accountId
	DB.Create(token)
	return token
}

func DeleteToken(accountId string, hash string) {
	DB.Where("hash = ? and account_id = ?", hash, accountId).Delete(&Token{})
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


