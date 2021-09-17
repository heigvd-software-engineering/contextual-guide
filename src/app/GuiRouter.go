package app

import (
	"github.com/gin-gonic/gin"
	"main/src/internal/controllers"
)

func initGuiRouter(router *gin.Engine) *gin.Engine {

	router.GET("/", controllers.Render)
	router.GET("/accounts/:accountId", controllers.GetAccount)

	router.GET("/resources", controllers.ListResources)
	router.GET("/resources/:id", controllers.ViewResource)
	router.GET("/resources/:id/qrcode.png", controllers.RenderResourceQRCode)
	router.GET("/resources/:id/redirect", controllers.RedirectResource)

	router.GET("/resources/create", controllers.RenderResourceForm)
	router.POST("/resources/create", controllers.CreateResource)

	router.GET("/tokens", controllers.GetTokens)
	router.GET("/tokens/create", controllers.RenderTokenForm)
	router.POST("/tokens/create", controllers.CreateToken)
	router.GET("/tokens/:id/delete", controllers.DeleteToken)

	return router

}
