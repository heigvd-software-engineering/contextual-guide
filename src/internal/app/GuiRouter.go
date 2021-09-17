package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"main/src/internal/controllers"
	"net/http"
	"os"
)

func initGuiRouter(router *gin.Engine) *gin.Engine {

	router.Use(extractCookie)

	router.GET("/", controllers.Render)

	router.GET("/tokens", checkLogged,  controllers.GetTokens)
	router.GET("/tokens/create", checkLogged,controllers.RenderTokenForm)
	router.POST("/tokens/create", checkLogged, controllers.CreateToken)
	router.GET("/tokens/:id/delete", checkLogged,controllers.DeleteToken)


	router.GET("/login", controllers.RenderLoginForm)
	router.POST("/login", controllers.HandleLogin)


	router.GET("/register", controllers.RenderRegisterForm)
	router.POST("/register", controllers.HandleRegistration)


	router.GET("/logout", controllers.HandleLogout)

	router.GET("/resources", controllers.ListResources)
	router.GET("/resources/:id", controllers.ViewResource)
	router.GET("/resources/:id/qrcode.png", controllers.RenderResourceQRCode)
	router.GET("/resources/:id/redirect", controllers.RedirectResource)

	router.GET("/resources/create", checkLogged,controllers.RenderResourceForm)
	router.POST("/resources/create", checkLogged, controllers.CreateResource)
	return router

}


func checkLogged(c *gin.Context)  {

	user, _ := c.Get("user")

	if user == nil {
		controllers.RenderErrorPage(http.StatusUnauthorized, "You are not authorized",c)
	}
}

func extractCookie(c *gin.Context) {
	jwtToken, err := c.Request.Cookie("sessionid")

	c.Set("user",nil)
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
			Id: accountId,
			Email: email,
		}
		c.Set("user",user)
	}
}