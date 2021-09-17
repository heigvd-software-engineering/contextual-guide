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

	router.GET("/", controllers.Render)
	router.GET("/accounts/:accountId", controllers.GetAccount)

	router.GET("/uris", controllers.GetUri)
	router.GET("/uris/:uuid", controllers.GetUriByUUID)

	router.GET("/uris/create", controllers.RenderUriForm)
	router.POST("/uris/create", controllers.CreateUri)

	router.GET("/tokens",checkAuthorization,  controllers.GetTokens)
	router.GET("/tokens/create", controllers.RenderTokenForm)
	router.POST("/tokens/create", controllers.CreateToken)
	router.GET("/tokens/:id/delete", controllers.DeleteToken)


	router.GET("/login", controllers.RenderLoginForm)
	router.POST("/login", controllers.HandleLogin)


	router.GET("/register", controllers.RenderRegisterForm)
	router.POST("/register", controllers.HandleRegistration)

	return router

}


func checkAuthorization(c *gin.Context) {
	jwtToken, err := c.Request.Cookie("sessionid")
	if err != nil {
		fmt.Println(err)
		controllers.RenderErrorPage(http.StatusUnauthorized, "You are not authorized",c)
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

		fmt.Println(accountId)

		c.Set("accountid",accountId)
		return

	}
	fmt.Println("YOP")
	controllers.RenderErrorPage(http.StatusUnauthorized, "You are not authorized",c)

}
