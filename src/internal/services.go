package internal

type accountService struct {
}

var (
	AccountService accountService
)

func (us *accountService) GetAccount(id string) *Account {
	user := AccountRepository.GetAccount(id)

	return user
}

func (us *accountService) CreateAccount(model *Account) (*Account, error)  {
	return AccountRepository.CreateAccount(model)
}

type resourceService struct {
}

var (
	ResourceService resourceService
)

func (us *resourceService) CreateResource(newResource *Resource) *Resource {
	resource := ResourceRepository.CreateResource(newResource)
	return resource
}

func (us *resourceService) GetAll() []Resource {
	resource := ResourceRepository.GetAllResource()
	return resource
}

func (us *resourceService) GetAllByAccountId(id string) []Resource {
	resource := ResourceRepository.GetAllResourceByAccountId(id)
	return resource
}

func (us *resourceService) GetOne(id string) *Resource {
	resource := ResourceRepository.GetResource(id)
	return resource
}

type tokenService struct{

}

var (
	TokenService tokenService
)

func NewToken() string {
	return ""
}

func (us *tokenService) CreateToken(newToken *Token) *Token {
	token := TokenRepository.CreateToken(newToken)
	return token
}
func (us *tokenService) GetAccountIdByToken(value string) string {
	token := TokenRepository.GetTokenByValue(value)

	if token == nil {
		return ""
	}

	return token.AccountId
}

func (us *tokenService) GetAllByAccountId(id string) []Token {
	token := TokenRepository.GetAllByAccountId(id)
	return token
}

func (us *tokenService) Delete(id int64) {
	TokenRepository.Delete(id)
}