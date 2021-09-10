package account

type IService interface {
	Create(*Model) bool
}


type service struct {
	accountRepository IRepository
}

func New(repository IRepository) IService {
	return &service{accountRepository: repository}
}

func (s *service) Create(accountToCreate *Model) bool  {
	s.accountRepository.Create(accountToCreate)

	return true
}
