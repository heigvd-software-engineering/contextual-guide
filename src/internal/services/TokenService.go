package services

import (
	"main/src/internal/models"
	"main/src/internal/repository"
)

type tokenService struct{

}

var (
	TokenService tokenService
)

func (us *tokenService) CreateToken(newToken *models.Token) *models.Token {
	token := repository.TokenRepository.CreateToken(newToken)
	return token
}
func (us *tokenService) GetAccountIdByToken(value string) string {
	token := repository.TokenRepository.GetTokenByValue(value)

	if token == nil {
		return ""
	}

	return token.AccountId
}

func (us *tokenService) GetAllByAccountId(id string) []models.Token {
	token := repository.TokenRepository.GetAllByAccountId(id)
	return token
}

func (us *tokenService) Delete(id int64) {
	repository.TokenRepository.Delete(id)
}
