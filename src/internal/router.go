package internal

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"main/src/internal/controllers"
	"main/src/internal/models"
)

var Engine *gin.Engine

func init() {
	models.ConnectDatabaseEnv()

	Engine = gin.Default()

	Engine.Static("/assets", "./src/assets")

	Engine = InitSite(Engine)
	Engine = InitApi(Engine)

	Engine.LoadHTMLGlob("src/internal/views/*/*.html")
	Engine.HTMLRender = createRender()

}

func InitSite(router *gin.Engine) *gin.Engine {
	router.Use(controllers.GetAccountFromCookie)
	router.GET("/", controllers.Render)
	router.GET("/tokens", controllers.IsAuthorized, controllers.GetTokens)
	router.GET("/tokens/create", controllers.IsAuthorized, controllers.RenderTokenForm)
	router.POST("/tokens/create", controllers.IsAuthorized, controllers.CreateToken)
	router.GET("/tokens/:hash/delete", controllers.IsAuthorized, controllers.DeleteToken)
	router.GET("/login", controllers.RenderLoginForm)
	router.POST("/login", controllers.HandleLogin)
	router.GET("/register", controllers.RenderRegisterForm)
	router.POST("/register", controllers.HandleRegistration)
	router.GET("/verify", controllers.RenderVerifyForm)
	router.POST("/verify", controllers.Verify)
	router.GET("/logout", controllers.HandleLogout)
	router.GET("/registry", controllers.LatestResources)
	router.GET("/resources", controllers.ListResources)
	router.GET("/resources/create", controllers.IsAuthorized, controllers.CreateResourceForm)
	router.POST("/resources", controllers.IsAuthorized, controllers.CreateResource)
	router.GET("/resources/:uuid", controllers.GetResource)
	router.GET("/resources/:uuid/update", controllers.UpdateResourceForm)
	router.POST("/resources/:uuid", controllers.UpdateResource)
	router.GET("/resources/:uuid/delete", controllers.DeleteResource)
	router.GET("/resources/:uuid/:size/qrcode.png", controllers.RenderResourceQRCode)
	router.GET("/resources/:uuid/redirect", controllers.RedirectResource)
	return router
}

func InitApi(router *gin.Engine) *gin.Engine {
	router.GET("/api/resources", controllers.GetAccountFromApiKey, controllers.IsAuthorized, controllers.ListResourceApi)
	router.GET("/api/resources/:uuid", controllers.GetAccountFromApiKey, controllers.IsAuthorized, controllers.GetResourceApi)
	router.POST("/api/resources", controllers.GetAccountFromApiKey, controllers.IsAuthorized, controllers.CreateResourceApi)
	router.PUT("/api/resources/:uuid", controllers.GetAccountFromApiKey, controllers.IsAuthorized, controllers.UpdateResourceApi)
	router.DELETE("/api/resources/:uuid", controllers.GetAccountFromApiKey, controllers.IsAuthorized, controllers.DeleteResourceApi)
	return router
}

func createRender() multitemplate.Renderer {
	renderer := multitemplate.NewRenderer()
	renderer.AddFromFiles("home", "src/internal/views/layouts/default.html", "src/internal/views/layouts/header.html", "src/internal/views/layouts/footer.html", "src/internal/views/pages/home.html")
	renderer.AddFromFiles("error", "src/internal/views/layouts/default.html", "src/internal/views/layouts/header.html", "src/internal/views/layouts/footer.html", "src/internal/views/pages/error.html")
	renderer.AddFromFiles("verify", "src/internal/views/layouts/default.html", "src/internal/views/layouts/header.html", "src/internal/views/layouts/footer.html", "src/internal/views/auth/validation-form.html")
	renderer.AddFromFiles("latest", "src/internal/views/layouts/default.html", "src/internal/views/layouts/header.html", "src/internal/views/layouts/footer.html", "src/internal/views/resources/latest.html")
	renderer.AddFromFiles("login-form", "src/internal/views/layouts/default.html", "src/internal/views/layouts/header.html", "src/internal/views/layouts/footer.html", "src/internal/views/auth/login-form.html")
	renderer.AddFromFiles("register-form", "src/internal/views/layouts/default.html", "src/internal/views/layouts/header.html", "src/internal/views/layouts/footer.html", "src/internal/views/auth/register-form.html")
	renderer.AddFromFiles("callback", "src/internal/views/layouts/default.html", "src/internal/views/layouts/header.html", "src/internal/views/layouts/footer.html", "src/internal/views/auth/callback.html")
	renderer.AddFromFiles("token-form", "src/internal/views/layouts/default.html", "src/internal/views/layouts/header.html", "src/internal/views/layouts/footer.html", "src/internal/views/tokens/token-form.html")
	renderer.AddFromFiles("token-list", "src/internal/views/layouts/default.html", "src/internal/views/layouts/header.html", "src/internal/views/layouts/footer.html", "src/internal/views/tokens/token-list.html")
	renderer.AddFromFiles("token-view", "src/internal/views/layouts/default.html", "src/internal/views/layouts/header.html", "src/internal/views/layouts/footer.html", "src/internal/views/tokens/token-view.html")
	renderer.AddFromFiles("resource-list", "src/internal/views/layouts/default.html", "src/internal/views/layouts/header.html", "src/internal/views/layouts/footer.html", "src/internal/views/resources/resource-list.html")
	renderer.AddFromFiles("resource-view", "src/internal/views/layouts/default.html", "src/internal/views/layouts/header.html", "src/internal/views/layouts/footer.html", "src/internal/views/resources/resource-view.html")
	renderer.AddFromFiles("resource-form", "src/internal/views/layouts/default.html", "src/internal/views/layouts/header.html", "src/internal/views/layouts/footer.html", "src/internal/views/resources/resource-form.html")
	return renderer
}
