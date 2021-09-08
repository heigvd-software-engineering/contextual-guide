package main

import (
	"database/sql"
	"fmt"
	"main/src/internal/account"
	"main/src/internal/storage"
	"net/http"
)

func home(w http.ResponseWriter, req *http.Request)  {
	fmt.Println("Hello")

	_, err := w.Write([]byte("Hello"))
	if err != nil {
		return 
	}
}


func Sum(x int, y int) int {
	return x+y
}

func setupDatabase(connString string) (*sql.DB, error) {
	// change "postgres" for whatever supported database you want to use
	db, err := sql.Open("postgres", connString)

	if err != nil {
		return nil, err
	}

	// ping the DB to ensure that it is connected
	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}

func run() error{
	connectionString := "postgres://postgres:postgres@localhost/**NAME-OF-YOUR-DATABASE-HERE**?sslmode=disable"

	// setup database connection
	db, err := setupDatabase(connectionString)

	if err != nil {
		return err
	}

	// create storage dependency
	datastore := storage.New(db)

	accountRepository := account.NewRepo(datastore)
	accountService := account.NewService(accountRepository)

	accountService.Patate()

	return nil

}

func main() {
	run()
	http.HandleFunc("/", home)
	_ = http.ListenAndServe(":3000", nil)
}