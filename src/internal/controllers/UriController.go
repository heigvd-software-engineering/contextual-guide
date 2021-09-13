package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"main/src/internal/models"
	"main/src/internal/services"
	"net/http"
)



func RenderUriForm(c *gin.Context) {
	c.HTML(http.StatusOK,"uri-form",gin.H{
		"title": "patate",
	})
}

func CreateUri(c *gin.Context)  {

	uri := models.Uri{
		Id: uuid.New().String(), Document: c.PostForm("document"),
	}

	services.UriService.CreateUri(&uri)

	c.Redirect(http.StatusOK,"/uris")
}

func GetUri(c *gin.Context) {

	uris := services.UriService.GetAll()

	c.HTML(http.StatusOK,"uri-list-view",gin.H{
		"uris": uris,
	})
}

func GetUriByUUID(c *gin.Context) {

	uriUUID := c.Param("uuid")
	fmt.Println(uriUUID)
	uri := services.UriService.GetOne(uriUUID)

	c.HTML(http.StatusOK,"uri-view",gin.H{
		"uri": uri,
	})
}
