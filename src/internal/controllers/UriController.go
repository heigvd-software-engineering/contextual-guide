package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	qrcode "github.com/skip2/go-qrcode"
	"log"
	"main/src/internal/models"
	"main/src/internal/services"
	"net/http"
)

func RenderUriForm(c *gin.Context) {
	c.HTML(http.StatusOK,"uri-form", nil)
}

func CreateUri(c *gin.Context) {

	uri := models.Uri{
		Id: uuid.New().String(), Document: c.PostForm("document"),
	}

	services.UriService.CreateUri(&uri)

	c.Redirect(http.StatusMovedPermanently,"/uris")
	c.Abort()

}

func GetUri(c *gin.Context) {

	uris := services.UriService.GetAll()

	c.HTML(http.StatusOK, "uri-list-view", gin.H{
		"uris": uris,
	})
}

func GetUriByUUID(c *gin.Context) {

	uriUUID := c.Param("uuid")
	uri := services.UriService.GetOne(uriUUID)

	c.HTML(http.StatusOK, "uri-view", gin.H{
		"uri": uri,
	})
}


func GetQRCode(c *gin.Context) {
	uriUUID := c.Param("uuid")
	uri := services.UriService.GetOne(uriUUID)

	// Unmarshal the json document
	blob := []byte(uri.Document)
	type Document struct {
		Redirect string
	}
	var document Document
	err := json.Unmarshal(blob, &document)
	if err != nil {
		log.Println(err)
		c.String(500, "Unable to unmarshal document")
		return
	}

	// Generate the QRCode
	png, err := qrcode.Encode(document.Redirect, qrcode.High, 256)
	if err != nil {
		log.Println(err)
		c.String(500, "Unable to encode qrcode")
		return
	}

	// Return a PNG file
	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Disposition", "attachment; filename=qrcode.png")
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Length", fmt.Sprintf("%d", len(png)))
	c.Writer.Write(png)
}
