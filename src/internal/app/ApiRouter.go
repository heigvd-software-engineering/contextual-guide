package app

import (
	"github.com/gin-gonic/gin"
	"main/src/internal/controllers"
)

func initApiRouter(router *gin.Engine) *gin.Engine {

	// scoped by the api-key
	router.GET("/api/resources",controllers.ListPrivateResourcesApi)

	router.POST("/api/resources",controllers.CreateResourceApi)

	router.GET("/api/resources/:id", controllers.ViewResourceApi)
	//router.PUT("/api/resources/:id", controllers.UpdateResourceApi)
	//router.DELETE("/api/resources/:id", controllers.ArchiveRessourceApi)

	return router
}
