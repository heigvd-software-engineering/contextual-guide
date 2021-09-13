package account

import (
	"log"
	"main/src/internal/storage"
	"main/src/pkg/httpserver"
)

type repository struct {
	datastore *storage.Storage
}

func NewRepo(datastore *storage.Storage) *httpserver.Repository {
	return &repository{datastore: datastore}
}



func (r *repository) Create(model *accountModel) (*httpserver.Model, bool){


	newAccountRequest := `INSERT INTO "account" (gotrueId) VALUES ($1);`

	err := r.datastore.Datastore.QueryRow(newAccountRequest, model.GoTrueId).Err()

	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return nil, true
	}

	return nil, false

}
func (r *repository) Update(id int64,model *httpserver.Model) (httpserver.Model, bool){
	return nil, false

}
func (r *repository) Delete(id int64) bool{

	return false

}
func (r *repository) Get()([]accountModel, bool){
	return nil, false
}

