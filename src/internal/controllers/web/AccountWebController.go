package webController

import (
	"main/src/internal"
)

func GetAccount(id string) *internal.Account {
	return internal.AccountService.GetAccount(id)
}
