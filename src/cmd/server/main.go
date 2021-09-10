package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"main/src/internal/account"
	"main/src/internal/storage"
	"net/http"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgresadmin"
	password = "admin123"
	dbname   = "postgresdb"
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

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)


	// setup database connection
	db, err := setupDatabase(psqlInfo)

	if err != nil {
		log.Fatal("Error connection to database", err)
	}

	// create storage dependency
	datastore := storage.New(db)

	accountRepository := account.NewRepo(datastore)
	accountService := account.New(accountRepository)

	account := account.Model{GoTrueId: "thisIsAnId"}
	accountService.Create(&account)

	return nil

}

func main() {
	run()
	http.HandleFunc("/", home)
	_ = http.ListenAndServe(":3000", nil)
}