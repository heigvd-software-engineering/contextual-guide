package account

import "main/src/pkg/httpserver"

type service struct {
	accountRepository *httpserver.Repository
}

func NewService(repository *httpserver.Repository) httpserver.Service {
	return &service{accountRepository: repository}
}

func (s *service) Create(accountToCreate *accountModel) bool  {
	s.accountRepository.C

	return true
}

func (s *service) GetAll() []httpserver.Model  {
	accounts, err := s.accountRepository.Get()

	if err {
		return nil
	}

	return accounts
}
