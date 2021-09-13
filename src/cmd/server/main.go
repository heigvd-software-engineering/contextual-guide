package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"main/src/internal/controllers"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}


func Sum(x int, y int) int {
	return x+y
}

func main() {
	port := flag.Int("port",3000, "-port=3000")
	flag.Parse()


	router.GET("/accounts/:accountId", controllers.GetAccount)
	router.POST("/accounts",controllers.CreateAccount)

	if err := router.Run(fmt.Sprintf(":%d",*port)); err != nil {
		panic(err)
	}
}