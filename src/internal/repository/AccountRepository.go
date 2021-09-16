package repository

import (
	"main/src/internal/database"
	"main/src/internal/models"
)

type accountRepository struct {
}

type IAccountRepository interface {
	GetAccount(string) *models.Account
	CreateAccount(*models.Account) (*models.Account, error)
}

var (
	AccountRepository IAccountRepository
)

func init() {
	AccountRepository = &accountRepository{}
}

func (ar *accountRepository) GetAccount(id string) *models.Account {
	var model models.Account
	database.DB.Where(&models.Account{ GoTrueId: id}).Find(&model)

	return &model
}

func (ar *accountRepository) CreateAccount(model *models.Account) (*models.Account, error) {
	err := database.DB.Create(model).Error

	if err != nil {
		return nil, err
	}
	return model, nil
}