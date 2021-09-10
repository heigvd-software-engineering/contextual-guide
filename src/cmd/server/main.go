package main

import (
	"flag"
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

func run() error{
	datastore := storage.New()

	accountRepository := account.NewRepo(datastore)
	accountService := account.New(accountRepository)

	account := account.Model{GoTrueId: "thisIsAnId"}
	accountService.Create(&account)

	return nil

}

func main() {
	port := flag.Int("port",3000, "-port=3000")
	flag.Parse()

	run();

	http.HandleFunc("/", home)

	_ = http.ListenAndServe(fmt.Sprintf(":%d",*port), nil)
}