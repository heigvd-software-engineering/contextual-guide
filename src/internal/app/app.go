package app

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"main/src/internal"
	"main/src/internal/controllers"
	"net/http"
)

var (
	Engine *gin.Engine
)

func createRender() multitemplate.Renderer {
	renderer := multitemplate.NewRenderer()
	renderer.AddFromFiles("home", "src/views/layouts/default.html", "src/views/layouts/header.html", "src/views/layouts/footer.html", "src/views/pages/home.html")
	renderer.AddFromFiles("error", "src/views/layouts/default.html", "src/views/layouts/header.html", "src/views/layouts/footer.html", "src/views/pages/error.html")
	renderer.AddFromFiles("verify", "src/views/layouts/default.html", "src/views/layouts/header.html", "src/views/layouts/footer.html", "src/views/auth/validation-form.html")
	renderer.AddFromFiles("resource-list-view-admin", "src/views/layouts/default.html", "src/views/layouts/header.html", "src/views/layouts/footer.html", "src/views/resources/resource-list-view-admin.html")
	renderer.AddFromFiles("resource-list-view", "src/views/layouts/default.html", "src/views/layouts/header.html", "src/views/layouts/footer.html", "src/views/resources/resource-list-view.html")
	renderer.AddFromFiles("resource-view", "src/views/layouts/default.html", "src/views/layouts/header.html", "src/views/layouts/footer.html", "src/views/resources/resource-view.html")
	renderer.AddFromFiles("resource-form", "src/views/layouts/default.html", "src/views/layouts/header.html", "src/views/layouts/footer.html", "src/views/resources/resource-form.html")

	renderer.AddFromFiles("token-form", "src/views/layouts/default.html", "src/views/layouts/header.html", "src/views/layouts/footer.html", "src/views/tokens/token-form.html")
	renderer.AddFromFiles("token-list-view-admin", "src/views/layouts/default.html", "src/views/layouts/header.html", "src/views/layouts/footer.html", "src/views/tokens/token-list-view-admin.html")
	renderer.AddFromFiles("created-token-view", "src/views/layouts/default.html", "src/views/layouts/header.html", "src/views/layouts/footer.html", "src/views/tokens/created-token-view.html")

	renderer.AddFromFiles("login-form", "src/views/layouts/default.html", "src/views/layouts/header.html", "src/views/layouts/footer.html", "src/views/auth/login-form.html")
	renderer.AddFromFiles("register-form", "src/views/layouts/default.html", "src/views/layouts/header.html", "src/views/layouts/footer.html", "src/views/auth/register-form.html")
	renderer.AddFromFiles("callback", "src/views/layouts/default.html", "src/views/layouts/header.html", "src/views/layouts/footer.html", "src/views/auth/callback.html")

	return renderer
}

func init() {
	internal.ConnectDatabase()

	Engine = gin.Default()

	Engine.Static("/assets", "./src/assets")

	Engine = initGuiRouter(Engine)
	Engine = initApiRouter(Engine)

	Engine.LoadHTMLGlob("src/views/*/*.html")
	Engine.HTMLRender = createRender()

}

func checkLogged(c *gin.Context)  {

	user, _ := c.Get("user")

	if user == nil {
		controllers.RenderErrorPage(http.StatusUnauthorized, "You are not authorized",c)
	}
}
