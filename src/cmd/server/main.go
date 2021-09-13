package main

import (
	"flag"
	"fmt"
	"main/src/internal/storage"
	"main/src/pkg/httpserver"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gin-gonic/gin"
)


import (
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}


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

func serve(port int){
	router.GET("/accounts/:accountId", controllers.GetUser)
	_ = http.ListenAndServe(fmt.Sprintf(":%d",port), router)

}

func main() {
	port := flag.Int("port",3000, "-port=3000")
	flag.Parse()
	datastore := storage.New()




	//accountModel := account.NewModel("ksdvkns")
	//
	//accountRepository := account.NewRepo(datastore)
	//accountService := account.NewService(accountRepository)
	//accountRouter := account.NewRouter(accountService)
	//
	//accountResource := httpserver.Resource{
	//	Name:  "Account",
	//	Model: accountModel,
	//	Router: accountRouter,
	//	Service: accountService,
	//	Repository: accountRepository,
	//}

	server := httpserver.New(*port,*datastore)



	server.Register(&accountResource)
	serve(*port)
}