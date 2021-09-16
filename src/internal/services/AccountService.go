package services

import (
	"main/src/internal/models"
	"main/src/internal/repository"
)

type accountService struct {
}

var (
	AccountService accountService
)

func (us *accountService) GetAccount(id string) *models.Account  {
	user := repository.AccountRepository.GetAccount(id)

	return user
}

func (us *accountService) CreateAccount(model *models.Account) (*models.Account, error)  {
	return repository.AccountRepository.CreateAccount(model)
}