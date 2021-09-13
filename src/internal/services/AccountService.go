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

func (us *accountService) GetAccount(accountId int64) *models.Account  {
	user := repository.AccountRepository.GetAccount(accountId)

	return user
}