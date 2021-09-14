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

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("home", "views/layouts/default.html","views/layouts/header.html", "views/layouts/footer.html","views/pages/home.html")
	r.AddFromFiles("uri-list-view", "views/layouts/default.html","views/layouts/header.html", "views/layouts/footer.html","views/uris/uri-list-view.html")
	r.AddFromFiles("uri-view", "views/layouts/default.html","views/layouts/header.html", "views/layouts/footer.html","views/uris/uri-view.html")
	r.AddFromFiles("uri-form", "views/layouts/default.html","views/layouts/header.html", "views/layouts/footer.html","views/uris/uri-form.html")

	r.AddFromFiles("token-form", "views/layouts/default.html","views/layouts/header.html", "views/layouts/footer.html","views/tokens/token-form.html")
	r.AddFromFiles("token-list-view", "views/layouts/default.html","views/layouts/header.html", "views/layouts/footer.html","views/tokens/token-list-view.html")
	r.AddFromFiles("created-token-view", "views/layouts/default.html","views/layouts/header.html", "views/layouts/footer.html","views/tokens/created-token-view.html")

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


	router.GET("/tokens",controllers.GetTokens)
	router.GET("/tokens/create",controllers.RenderTokenForm)
	router.POST("/tokens/create",controllers.CreateToken)
	router.GET("/tokens/:id/delete",controllers.DeleteToken)

	if err := router.Run(fmt.Sprintf(":%d",*port)); err != nil {
		panic(err)
	}
}