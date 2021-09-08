package account

import "main/src/internal/storage"

type Model struct {
	GoTrueId string
}

type IService interface {
	Patate() bool
}

type IRepository interface {
	Create(*Model) Model
	Update(int64,*Model) Model
	Delete(int64)
	Get()[]Model
}

type service struct {
	accountRepository IRepository
}
type repository struct {
	datastore storage.Storage
}

func NewRepo(datastore storage.Storage) IRepository {
	return &repository{datastore: datastore}
}


func NewService(repository IRepository) IService {
	return &service{accountRepository: repository}
}

func (r *repository) Create(model *Model) Model{
	r.datastore.CreateAccount(model)
	return Model{}

}
func (r *repository) Update(id int64,model *Model) Model{
	return Model{}

}
func (r *repository) Delete(id int64){

}
func (r *repository) Get()[]Model{
	return nil
}

func (s *service) Patate() bool  {
	return true
}
