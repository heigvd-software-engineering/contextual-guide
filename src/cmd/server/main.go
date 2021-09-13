package main

import (
	"flag"
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"main/src/internal/controllers"
)

var (
	router *gin.Engine
)



func init() {
	router = gin.Default()
	router.LoadHTMLGlob("views/*.html")
	router.HTMLRender = createMyRender()
}


func Sum(x int, y int) int {
	return x+y
}


func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("home", "views/layouts/default.html", "views/pages/home.html")
	return r
}



func main() {
	port := flag.Int("port",3000, "-port=3000")
	flag.Parse()

	router.GET("/", controllers.Render)

	router.GET("/accounts/:accountId", controllers.GetAccount)
	router.GET("/uri/create",controllers.RenderUriForm)
	router.POST("/uri/create",controllers.CreateUri)
	router.GET("/uri",controllers.GetUri)

	if err := router.Run(fmt.Sprintf(":%d",*port)); err != nil {
		panic(err)
	}
}