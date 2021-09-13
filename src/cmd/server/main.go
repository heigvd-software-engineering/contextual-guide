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
	router.LoadHTMLGlob("views/*/*.html")
	router.HTMLRender = createMyRender()
}


func Sum(x int, y int) int {
	return x+y
}


func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("home", "views/layouts/default.html","views/layouts/header.html", "views/layouts/footer.html","views/pages/home.html")
	r.AddFromFiles("uri-list-view", "views/layouts/default.html","views/layouts/header.html", "views/layouts/footer.html","views/uris/uri-list-view.html")
	r.AddFromFiles("uri-view", "views/layouts/default.html","views/layouts/header.html", "views/layouts/footer.html","views/uris/uri-view.html")
	r.AddFromFiles("uri-form", "views/layouts/default.html","views/layouts/header.html", "views/layouts/footer.html","views/uris/uri-form.html")

	return r
}



func main() {
	port := flag.Int("port",3000, "-port=3000")
	flag.Parse()

	router.Static("/assets", "./assets")

	router.GET("/", controllers.Render)
	router.GET("/accounts/:accountId", controllers.GetAccount)

	router.GET("/uris",controllers.GetUri)
	router.GET("/uris/:uuid",controllers.GetUriByUUID)

	router.GET("/uris/create",controllers.RenderUriForm)
	router.POST("/uris/create",controllers.CreateUri)

	if err := router.Run(fmt.Sprintf(":%d",*port)); err != nil {
		panic(err)
	}
}