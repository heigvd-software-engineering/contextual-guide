package services

import (
	"main/src/internal/database"
	"main/src/internal/models"
)

func GetToken(id int64) *models.Token {
	var tokens models.Token
	//database.DB.Where(&models.Token{: id}).Find(&tokens)
	return &tokens
}

func CreateToken(model *models.Token) *models.Token {
	database.DB.Create(model)
	return model
}

func ListTokenByAccountId(id string) []models.Token {
	tokens := []models.Token{}
	database.DB.Where(&models.Token{AccountId: id}).Find(&tokens)
	return tokens
}

func DeleteToken(id int64) {
	database.DB.Delete(&models.Token{}, id)
}

func GetTokenByValue(value string) *models.Token {
	var tokens models.Token
	database.DB.Where(&models.Token{Hash: value}).Find(&tokens)
	return &tokens
}
