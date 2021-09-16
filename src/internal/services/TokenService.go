package services

import (
	"main/src/internal/models"
	"main/src/internal/repository"
)

type tokenService struct {
}

var (
	TokenService tokenService
)

func (us *tokenService) CreateToken(newToken *models.Token) *models.Token {

	token := repository.TokenRepository.CreateToken(newToken)
	return token
}

func (us *tokenService) GetAll() []models.Token {

	token := repository.TokenRepository.GetAll()
	return token
}

func (us *tokenService) Delete(id int64) {
	repository.TokenRepository.Delete(id)
}
