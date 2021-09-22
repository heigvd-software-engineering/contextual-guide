package webController

import (
	"main/src/internal/models"
	"main/src/internal/services"
)

func GetAccount(id string) *models.Account {
	return services.AccountService.GetAccount(id)
}
