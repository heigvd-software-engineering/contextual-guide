package storage

import (
	"database/sql"
	"log"
	"main/src/internal/account"
)

type Storage interface{
	CreateAccount(model *account.Model) error

}

type storage struct {
	db *sql.DB
}

func (s *storage) CreateAccount(model *account.Model) error{

	newAccountRequest := `INSERT INTO "account" (gotrueId) VALUES ($1);`

	err := s.db.QueryRow(newAccountRequest, model.GoTrueId).Err()

	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return err
	}

	return nil
}

func New(db *sql.DB) Storage {
	return &storage{
		db: db,
	}
}
