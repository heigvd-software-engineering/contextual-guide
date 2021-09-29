package internal

type accountRepository struct {
}

type IAccountRepository interface {
	GetAccount(string) *Account
	CreateAccount(*Account) (*Account, error)
}

var (
	AccountRepository IAccountRepository
)

func init() {
	AccountRepository = &accountRepository{}
}

func (ar *accountRepository) GetAccount(id string) *Account {
	var model Account
	DB.Where(&Account{GoTrueId: id}).Find(&model)

	return &model
}

func (ar *accountRepository) CreateAccount(model *Account) (*Account, error) {
	err := DB.Create(model).Error

	if err != nil {
		return nil, err
	}
	return model, nil
}

type tokenRepository struct {
}

type ITokenRepository interface {
	GetToken(int64) *Token
	Delete(int64)
	GetAllByAccountId(string) []Token
	CreateToken(*Token) *Token
	GetTokenByValue(string) *Token
}

var (
	TokenRepository ITokenRepository
)

func init() {
	TokenRepository = &tokenRepository{}
}

func (ur *tokenRepository) GetToken(id int64) *Token {
	var tokens Token
	//database.DB.Where(&models.Token{: id}).Find(&tokens)
	return &tokens
}

func (ur *tokenRepository) CreateToken(model *Token) *Token {
	DB.Create(model)

	return model
}

func (ur *tokenRepository) GetAllByAccountId(id string) []Token {
	tokens := []Token{}

	DB.Where(&Token{AccountId: id}).Find(&tokens)
	return tokens
}

func (ur *tokenRepository) Delete(id int64) {
	DB.Delete(&Token{}, id)
}

func (ur *tokenRepository) GetTokenByValue(value string) *Token {
	var tokens Token
	DB.Where(&Token{Hash: value}).Find(&tokens)
	return &tokens
}

type resourceRepository struct {
}

type IResourceRepository interface {
	GetResource(string) *Resource
	GetAllResource() []Resource
	GetAllResourceByAccountId(string) []Resource
	CreateResource(*Resource) *Resource
}

var (
	ResourceRepository IResourceRepository
)

func init() {
	ResourceRepository = &resourceRepository{}
}

func (ur *resourceRepository) GetResource(id string) *Resource {
	var resource Resource
	DB.Where(&Resource{Uuid: id}).Find(&resource)
	return &resource
}

func (ur *resourceRepository) CreateResource(model *Resource) *Resource {
	DB.Create(model)
	return model
}

func (ur *resourceRepository) GetAllResource() []Resource {
	var resources []Resource
	DB.Preload("Account", &resources)
	return resources
}

func (ur *resourceRepository) GetAllResourceByAccountId(id string) []Resource {
	var resources []Resource
	DB.Preload("Account", &resources)
	DB.Where(&Resource{AccountId: id}).Find(&resources)
	return resources
}
