package app

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

var (
	Engine *gin.Engine
)


func createRender() multitemplate.Renderer {
	renderer := multitemplate.NewRenderer()
	renderer.AddFromFiles("home", "views/layouts/default.html","views/layouts/header.html", "views/layouts/footer.html","views/pages/home.html")
	renderer.AddFromFiles("uri-list-view", "views/layouts/default.html","views/layouts/header.html", "views/layouts/footer.html","views/uris/uri-list-view.html")
	renderer.AddFromFiles("uri-view", "views/layouts/default.html","views/layouts/header.html", "views/layouts/footer.html","views/uris/uri-view.html")
	renderer.AddFromFiles("uri-form", "views/layouts/default.html","views/layouts/header.html", "views/layouts/footer.html","views/uris/uri-form.html")

	renderer.AddFromFiles("token-form", "views/layouts/default.html","views/layouts/header.html", "views/layouts/footer.html","views/tokens/token-form.html")
	renderer.AddFromFiles("token-list-view", "views/layouts/default.html","views/layouts/header.html", "views/layouts/footer.html","views/tokens/token-list-view.html")
	renderer.AddFromFiles("created-token-view", "views/layouts/default.html","views/layouts/header.html", "views/layouts/footer.html","views/tokens/created-token-view.html")

	return renderer
}


func init(){

	Engine = gin.Default()

	Engine.Static("/assets", "./assets")

	Engine = initGuiRouter(Engine)
	Engine = initApiRouter(Engine)

	Engine.LoadHTMLGlob("views/*/*.html")
	Engine.HTMLRender = createRender()

}
