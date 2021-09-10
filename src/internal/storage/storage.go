package storage

import (
	"database/sql"
)

type Storage struct {
	Datastore *sql.DB
}

func New(db *sql.DB) *Storage {
	return &Storage{
		Datastore: db,
	}
}
