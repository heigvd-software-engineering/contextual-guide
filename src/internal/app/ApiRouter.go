package app

import (
	"github.com/gin-gonic/gin"
	"main/src/internal"
	"main/src/internal/controllers"
	apiController "main/src/internal/controllers/api"
	"net/http"
)

func initApiRouter(router *gin.Engine) *gin.Engine {

	// scoped by the api-key
	router.GET("/api/resources",getAccountFromApiKey, checkLogged, apiController.ListPrivateResources)
	router.POST("/api/resources",getAccountFromApiKey, checkLogged, apiController.CreateResource)
	router.GET("/api/resources/:id",getAccountFromApiKey, checkLogged, apiController.ViewResource)

	//router.PUT("/api/resources/:id", controllers.UpdateResourceApi)
	//router.DELETE("/api/resources/:id", controllers.ArchiveRessourceApi)

	return router
}

func getAccountFromApiKey(c *gin.Context)  {
	c.Set("user",nil)

	key := c.Request.Header.Get("x-api-key")
	if key == "" {
		c.JSON(http.StatusUnauthorized,"You are not authorized")
	}

	accountId := internal.TokenService.GetAccountIdByToken(key)
	if accountId == "" {
		c.JSON(http.StatusUnauthorized,"You are not authorized")
	}

	user := controllers.LoggedUser{
		Id: accountId,
		Email: "",
	}

	c.Set("user",user)
}