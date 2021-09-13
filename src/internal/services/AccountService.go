package services

import "main/src/internal/models"

type AccountService struct {

}

var (
	accountService AccountService
)

func (us *AccountService) GetAccount(accountId int64) (*models.Account)  {
	user, err := domain.UserDao.GetUser(userID)

	if err != nil {
		return nil
	}

	return user
}