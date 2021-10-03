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
	renderer.AddFromFiles("home", "src/internal/views/layouts/default.html", "src/internal/views/layouts/header.html", "src/internal/views/layouts/footer.html", "src/internal/views/pages/home.html")
	renderer.AddFromFiles("error", "src/internal/views/layouts/default.html", "src/internal/views/layouts/header.html", "src/internal/views/layouts/footer.html", "src/internal/views/pages/error.html")
	renderer.AddFromFiles("verify", "src/internal/views/layouts/default.html", "src/internal/views/layouts/header.html", "src/internal/views/layouts/footer.html", "src/internal/views/auth/validation-form.html")
	renderer.AddFromFiles("resource-list-view-admin", "src/internal/views/layouts/default.html", "src/internal/views/layouts/header.html", "src/internal/views/layouts/footer.html", "src/internal/views/resources/resource-list.html")
	renderer.AddFromFiles("resource-list-view", "src/internal/views/layouts/default.html", "src/internal/views/layouts/header.html", "src/internal/views/layouts/footer.html", "src/internal/views/resources/registry.html")
	renderer.AddFromFiles("resource-view", "src/internal/views/layouts/default.html", "src/internal/views/layouts/header.html", "src/internal/views/layouts/footer.html", "src/internal/views/resources/resource-view.html")
	renderer.AddFromFiles("resource-form", "src/internal/views/layouts/default.html", "src/internal/views/layouts/header.html", "src/internal/views/layouts/footer.html", "src/internal/views/resources/resource-form.html")

	renderer.AddFromFiles("token-form", "src/internal/views/layouts/default.html", "src/internal/views/layouts/header.html", "src/internal/views/layouts/footer.html", "src/internal/views/tokens/token-form.html")
	renderer.AddFromFiles("token-list-view-admin", "src/internal/views/layouts/default.html", "src/internal/views/layouts/header.html", "src/internal/views/layouts/footer.html", "src/internal/views/tokens/token-list.html")
	renderer.AddFromFiles("created-token-view", "src/internal/views/layouts/default.html", "src/internal/views/layouts/header.html", "src/internal/views/layouts/footer.html", "src/internal/views/tokens/token-view.html")

	renderer.AddFromFiles("login-form", "src/internal/views/layouts/default.html", "src/internal/views/layouts/header.html", "src/internal/views/layouts/footer.html", "src/internal/views/auth/login-form.html")
	renderer.AddFromFiles("register-form", "src/internal/views/layouts/default.html", "src/internal/views/layouts/header.html", "src/internal/views/layouts/footer.html", "src/internal/views/auth/register-form.html")
	renderer.AddFromFiles("callback", "src/internal/views/layouts/default.html", "src/internal/views/layouts/header.html", "src/internal/views/layouts/footer.html", "src/internal/views/auth/callback.html")

	return renderer
}

func init() 	{
	models.ConnectDatabaseEnv()

	Engine = gin.Default()

	Engine.Static("/assets", "./src/assets")

	Engine = addSiteRoutes(Engine)
	Engine = initApiRouter(Engine)

	Engine.LoadHTMLGlob("src/internal/views/*/*.html")
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
	router.GET("/tokens/:hash/delete", checkLogged, controllers.DeleteToken)

	router.GET("/login", controllers.RenderLoginForm)
	router.POST("/login", controllers.HandleLogin)

	router.GET("/register", controllers.RenderRegisterForm)
	router.POST("/register", controllers.HandleRegistration)

	router.GET("/verify", controllers.RenderVerifyForm)
	router.POST("/verify", controllers.Verify)

	router.GET("/logout", controllers.HandleLogout)

	router.GET("/registry", controllers.Registry)

	router.GET("/resources", controllers.ListResources)
	router.GET("/resources/create", checkLogged, controllers.ResourceForm)
	router.POST("/resources", checkLogged, controllers.CreateResource)
	router.GET("/resources/:uuid", controllers.ReadResource)
	router.GET("/resources/:uuid/update", controllers.ResourceForm)
	router.POST("/resources/:uuid", controllers.UpdateResource)
	router.GET("/resources/:uuid/delete", controllers.DeleteResource)


	router.GET("/resources/:uuid/:size/qrcode.png", controllers.RenderResourceQRCode)
	router.GET("/resources/:uuid/redirect", controllers.RedirectResource)


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

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token != nil && token.Valid {
		accountId, _ := claims["sub"].(string)
		email, _ := claims["email"].(string)
		user := controllers.User{
			Id:    accountId,
			Email: email,
		}
		c.Set("user", user)
	}
}

// API
func initApiRouter(router *gin.Engine) *gin.Engine {

	/*
	// scoped by the api-key
	router.GET("/api/resources", getAccountFromApiKey, checkLogged, controllers.GetResources)
	router.POST("/api/resources", getAccountFromApiKey, checkLogged, controllers.PostResource)
	router.GET("/api/resources/:id", getAccountFromApiKey, checkLogged, controllers.GetResource)

	//router.PUT("/api/resources/:id", controllers.UpdateResourceApi)
	//router.DELETE("/api/resources/:id", controllers.ArchiveRessourceApi)
	 */
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

	user := controllers.User{
		Id:    token.AccountId,
		Email: "",
	}

	c.Set("user", user)
}
