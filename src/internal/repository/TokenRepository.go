package repository

import (
	"main/src/internal/database"
	"main/src/internal/models"
)

type tokenRepository struct {
}

type ITokenRepository interface {
	GetToken(int64) *models.Token
	Delete(int64)
	GetAllByAccountId(string) []models.Token
	CreateToken(*models.Token) *models.Token
	GetTokenByValue(string) *models.Token
}

var (
	TokenRepository ITokenRepository
)

func init() {
	TokenRepository = &tokenRepository{}
}

func (ur *tokenRepository) GetToken(id int64) *models.Token {
	var tokens models.Token
	//database.DB.Where(&models.Token{: id}).Find(&tokens)
	return &tokens
}

func (ur *tokenRepository) CreateToken(model *models.Token) *models.Token {
	database.DB.Create(model)

	return model
}

func (ur *tokenRepository) GetAllByAccountId(id string) []models.Token {
	tokens := []models.Token{}

	database.DB.Where(&models.Token{AccountId: id}).Find(&tokens)
	return tokens
}

func (ur *tokenRepository) Delete(id int64) {
	database.DB.Delete(&models.Token{}, id)
}

func (ur *tokenRepository) GetTokenByValue(value string) *models.Token {
	var tokens models.Token
	database.DB.Where(&models.Token{Value: value}).Find(&tokens)

	return &tokens
}