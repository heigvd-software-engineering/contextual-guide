package app

import (
	"github.com/gin-gonic/gin"
	"main/src/internal/controllers"
)

func initGuiRouter(router *gin.Engine) *gin.Engine {

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

	router.GET("/register",controllers.RenderRegisterForm)
	router.POST("/register",controllers.HandleRegistration)

	router.GET("/login",controllers.RenderLoginForm)
	router.POST("/login",controllers.HandleLogin)



	return router

}

/*func checkAuthorisation(c *gin.Context) {
	token, err := c.Request.Cookie("sessionid")
	if err != nil {
		c.Redirect(http.StatusUnauthorized, "/")
	}

	c.Keys

}*/
