package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type connectionInfo struct {
	host string
	port string
	user string
	password string
	dbname string
}

type Storage struct {
	Datastore *sql.DB
}


func connect(connectionInfo *connectionInfo) (*sql.DB, error) {

	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		connectionInfo.host, connectionInfo.port, connectionInfo.user, connectionInfo.password, connectionInfo.dbname)

	fmt.Println(connString)
	db, err := sql.Open("postgres", connString)

	if err != nil {
		return nil, err
	}
	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}

func New() *Storage {

	databaseConnectionInfo := connectionInfo{
		port: os.Getenv("DB_PORT"),
		host: os.Getenv("DB_HOST"),
		user: os.Getenv("DB_USER"),
		dbname: os.Getenv("DB_NAME"),
		password: os.Getenv("DB_PASS"),
	}


	db, err := connect(&databaseConnectionInfo)

	if err != nil {
		log.Fatal("Error when connecting to database :", err)
	}

	return &Storage{
		Datastore: db,
	}
}
