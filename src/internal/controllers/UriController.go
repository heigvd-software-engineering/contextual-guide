package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"main/src/internal/models"
	"main/src/internal/services"
	"net/http"
)



func RenderUriForm(c *gin.Context) {
	c.HTML(http.StatusOK,"uri-form.html",gin.H{
		"title": "patate",
	})
}

func CreateUri(c *gin.Context)  {

	uri := models.Uri{
		Id: uuid.New().String(), Document: c.PostForm("document"),
	}

	services.UriService.CreateUri(&uri)

	c.Redirect(http.StatusOK,"/uri")
}

func GetUri(c *gin.Context) {

	uris := services.UriService.GetAll()

	c.HTML(http.StatusOK,"uri-view.html",gin.H{
		"uris": uris,
	})
}
