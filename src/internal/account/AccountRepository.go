package account

import (
	"log"
	"main/src/internal/storage"
)

type IRepository interface {
	Create(*Model) (Model,bool)
	Update(int64,*Model) (Model, bool)
	Delete(int64) bool
	Get()([]Model, bool)
}

type repository struct {
	datastore *storage.Storage
}

func NewRepo(datastore *storage.Storage) IRepository {
	return &repository{datastore: datastore}

}



func (r *repository) Create(model *Model) (Model, bool){


	newAccountRequest := `INSERT INTO "account" (gotrueId) VALUES ($1);`

	err := r.datastore.Datastore.QueryRow(newAccountRequest, model.GoTrueId).Err()

	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return Model{}, true
	}

	return Model{}, false

}
func (r *repository) Update(id int64,model *Model) (Model, bool){
	return Model{}, false

}
func (r *repository) Delete(id int64) bool{

	return false

}
func (r *repository) Get()([]Model, bool){
	return nil, false
}

