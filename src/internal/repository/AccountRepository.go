package repository

import "main/src/internal/models"

type accountRepository struct {
}

type IAccountRepository interface {
	GetAccount(int64) *models.Account
}

var (
	AccountRepository IAccountRepository
)

func init() {
	AccountRepository = &accountRepository{}
}

func (ar *accountRepository) GetAccount(id int64) *models.Account {
	return &models.Account{GoTrueId: "test"}
}
