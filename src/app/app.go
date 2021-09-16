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
	renderer.AddFromFiles("home", "src/views/layouts/default.html", "src/views/layouts/header.html", "src/views/layouts/footer.html", "src/views/pages/home.html")
	renderer.AddFromFiles("uri-list-view", "src/views/layouts/default.html", "src/views/layouts/header.html", "src/views/layouts/footer.html", "src/views/uris/uri-list-view.html")
	renderer.AddFromFiles("uri-view", "src/views/layouts/default.html", "src/views/layouts/header.html", "src/views/layouts/footer.html", "src/views/uris/uri-view.html")
	renderer.AddFromFiles("uri-form", "src/views/layouts/default.html", "src/views/layouts/header.html", "src/views/layouts/footer.html", "src/views/uris/uri-form.html")

	renderer.AddFromFiles("token-form", "src/views/layouts/default.html", "src/views/layouts/header.html", "src/views/layouts/footer.html", "src/views/tokens/token-form.html")
	renderer.AddFromFiles("token-list-view", "src/views/layouts/default.html", "src/views/layouts/header.html", "src/views/layouts/footer.html", "src/views/tokens/token-list-view.html")
	renderer.AddFromFiles("created-token-view", "src/views/layouts/default.html", "src/views/layouts/header.html", "src/views/layouts/footer.html", "src/views/tokens/created-token-view.html")

	return renderer
}

func init() {

	Engine = gin.Default()

	Engine.Static("/assets", "./src/assets")

	Engine = initGuiRouter(Engine)
	Engine = initApiRouter(Engine)

	Engine.LoadHTMLGlob("src/views/*/*.html")
	Engine.HTMLRender = createRender()

}
