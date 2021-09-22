package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"main/src/internal/controllers"
	"main/src/internal/controllers/web"
	"os"
)

func initGuiRouter(router *gin.Engine) *gin.Engine {

	router.Use(extractCookie)

	router.GET("/", controllers.Render)

	// Get account related token
	router.GET("/tokens", checkLogged, webController.GetTokens)
	router.GET("/tokens/create", checkLogged, webController.RenderTokenForm)
	router.POST("/tokens/create", checkLogged, webController.CreateToken)
	router.GET("/tokens/:id/delete", checkLogged, webController.DeleteToken)


	router.GET("/login", webController.RenderLoginForm)
	router.POST("/login", webController.HandleLogin)


	router.GET("/register", webController.RenderRegisterForm)
	router.POST("/register", webController.HandleRegistration)

	router.GET("/verify", webController.RenderVerifyForm)
	router.POST("/verify", webController.Verfify)

	router.GET("/logout", webController.HandleLogout)

	router.GET("/resources", webController.ListAllResources)
	router.GET("/resources/mine", webController.ListPrivateResources)
	router.GET("/resources/:id", webController.ViewResource)
	router.GET("/resources/:id/qrcode.png", webController.RenderResourceQRCode)
	router.GET("/resources/:id/redirect", webController.RedirectResource)

	router.GET("/resources/create", checkLogged, webController.RenderResourceForm)
	router.POST("/resources/create", checkLogged, webController.CreateResource)
	return router

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
