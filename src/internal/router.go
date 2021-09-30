package internal

import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"main/src/internal/controllers"
	"main/src/internal/models"
	"net/http"
	"os"
)

var Engine *gin.Engine

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

func init() 	{
	models.ConnectDatabase()

	Engine = gin.Default()

	Engine.Static("/assets", "./src/assets")

	Engine = addSiteRoutes(Engine)
	Engine = initApiRouter(Engine)

	Engine.LoadHTMLGlob("src/views/*/*.html")
	Engine.HTMLRender = createRender()

}

func checkLogged(c *gin.Context) {
	user, _ := c.Get("user")
	if user == nil {
		controllers.RenderErrorPage(http.StatusUnauthorized, "You are not authorized", c)
	}
}

func addSiteRoutes(router *gin.Engine) *gin.Engine {

	router.Use(extractCookie)

	router.GET("/", controllers.Render)

	// Get account related token
	router.GET("/tokens", checkLogged, controllers.GetTokens)
	router.GET("/tokens/create", checkLogged, controllers.RenderTokenForm)
	router.POST("/tokens/create", checkLogged, controllers.CreateToken)
	router.GET("/tokens/:id/delete", checkLogged, controllers.DeleteToken)

	router.GET("/login", controllers.RenderLoginForm)
	router.POST("/login", controllers.HandleLogin)

	router.GET("/register", controllers.RenderRegisterForm)
	router.POST("/register", controllers.HandleRegistration)

	router.GET("/verify", controllers.RenderVerifyForm)
	router.POST("/verify", controllers.Verify)

	router.GET("/logout", controllers.HandleLogout)

	router.GET("/registry", controllers.Registry)
	router.GET("/resources", controllers.ListResources)
	router.GET("/resources/:id", controllers.ViewResource)
	router.GET("/resources/:id/qrcode.png", controllers.RenderResourceQRCode)
	router.GET("/resources/:id/redirect", controllers.RedirectResource)

	router.GET("/resources/create", checkLogged, controllers.RenderResourceForm)
	router.POST("/resources/create", checkLogged, controllers.CreateResource)
	return router

}

func extractCookie(c *gin.Context) {
	jwtToken, err := c.Request.Cookie("sessionid")

	c.Set("user", nil)
	if err != nil {
		return
	}

	secret := os.Getenv("JWT_SECRET")
	token, _ := jwt.Parse(jwtToken.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		accountId, _ := claims["sub"].(string)
		email, _ := claims["email"].(string)

		user := controllers.LoggedUser{
			Id:    accountId,
			Email: email,
		}
		c.Set("user", user)
	}
}

// API
func initApiRouter(router *gin.Engine) *gin.Engine {

	// scoped by the api-key
	router.GET("/api/resources", getAccountFromApiKey, checkLogged, controllers.GetResources)
	router.POST("/api/resources", getAccountFromApiKey, checkLogged, controllers.PostResource)
	router.GET("/api/resources/:id", getAccountFromApiKey, checkLogged, controllers.GetResource)

	//router.PUT("/api/resources/:id", controllers.UpdateResourceApi)
	//router.DELETE("/api/resources/:id", controllers.ArchiveRessourceApi)

	return router
}

func getAccountFromApiKey(c *gin.Context) {
	c.Set("user", nil)

	key := c.Request.Header.Get("x-api-key")
	if key == "" {
		c.JSON(http.StatusUnauthorized, "You are not authorized")
	}

	token := models.GetTokenByValue(key)
	if token == nil {
		c.JSON(http.StatusUnauthorized, "You are not authorized")
	}

	user := controllers.LoggedUser{
		Id:    token.AccountId,
		Email: "",
	}

	c.Set("user", user)
}
