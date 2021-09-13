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
	router.LoadHTMLGlob("view/*.html")
}


func Sum(x int, y int) int {
	return x+y
}




func main() {
	port := flag.Int("port",3000, "-port=3000")
	flag.Parse()


	router.GET("/accounts/:accountId", controllers.GetAccount)
	router.GET("/uri/create",controllers.RenderUriForm)
	router.POST("/uri/create",controllers.CreateUri)
	router.GET("/uri",controllers.GetUri)

	if err := router.Run(fmt.Sprintf(":%d",*port)); err != nil {
		panic(err)
	}
}