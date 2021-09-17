package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v3"
	qrcode "github.com/skip2/go-qrcode"
	"log"
	"main/src/internal/models"
	"main/src/internal/services"
	"net/http"
)

type Content struct {
	Redirect string
}

func RenderResourceForm(c *gin.Context) {
	c.HTML(http.StatusOK, "resource-form", nil)
}

func CreateResource(c *gin.Context) {
	resource := models.Resource{
		Id: shortuuid.New(), Content: c.PostForm("resource"),
	}

	services.ResourceService.CreateResource(&resource)

	c.Redirect(http.StatusMovedPermanently, "/resources")
	c.Abort()
}

func ListResources(c *gin.Context) {
	resources := services.ResourceService.GetAll()

	c.HTML(http.StatusOK, "resource-list-view", gin.H{
		"resources": resources,
	})
}

func ViewResource(c *gin.Context) {
	resourceId := c.Param("id")
	resource := services.ResourceService.GetOne(resourceId)

	c.HTML(http.StatusOK, "resource-view", gin.H{
		"resource": resource,
	})
}

func RenderResourceQRCode(c *gin.Context) {
	resourceId := c.Param("id")

	// Generate the QRCode
	uri := fmt.Sprintf("/resources/%s", resourceId)
	png, err := qrcode.Encode(uri, qrcode.High, 256)
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

func RedirectResource(c *gin.Context) {
	resourceId := c.Param("id")
	resource := services.ResourceService.GetOne(resourceId)

	// Unmarshal the resource content
	blob := []byte(resource.Content)
	var content Content
	err := json.Unmarshal(blob, &resource)
	if err != nil {
		log.Println(err)
		c.String(500, "Unable to unmarshal resource")
		return
	}

	if content.Redirect != "" {
		c.Redirect(http.StatusMovedPermanently, content.Redirect)
	} else {
		uri := fmt.Sprintf("/resources/%s", resourceId)
		c.Redirect(http.StatusMovedPermanently, uri)
	}

	c.Abort()

}
