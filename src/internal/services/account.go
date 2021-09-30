package services

import (
	"main/src/internal/database"
	"main/src/internal/models"
)

func GetAccount(id string) *models.Account {
	var model models.Account
	database.DB.Where(&models.Account{GoTrueId: id}).Find(&model)
	return &model
}

func CreateAccount(model *models.Account) (*models.Account, error) {
	err := database.DB.Create(model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}
