package webController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v3"
	qrcode "github.com/skip2/go-qrcode"
	"log"
	"main/src/internal/controllers"
	"main/src/internal/models"
	"main/src/internal/services"
	"net/http"
	"strconv"
	"time"
)


func RenderResourceForm(c *gin.Context) {
	c.HTML(http.StatusOK, "resource-form", gin.H{
		"user": controllers.GetUserFromContext(c),
	})
}

func CreateResource(c *gin.Context) {

	account := services.AccountService.GetAccount(controllers.GetUserFromContext(c).Id)

	longitude, _ := strconv.ParseFloat(c.PostForm("longitude"), 32)
	latitude, _ := strconv.ParseFloat(c.PostForm("latitude"), 32)

	timeLayout := "2006-01-02T15:04:05.000Z"
	timestamp, _ := time.Parse(timeLayout, c.PostForm("timestamp"))

	resource := models.Resource{
		Uuid:        shortuuid.New(),
		Title:       c.PostForm("title"),
		Description: c.PostForm("description"),
		Timestamp:   timestamp,
		Longitude:   float32(longitude),
		Latitude:    float32(latitude),
		Redirect:    c.PostForm("redirect"),
		AccountId:   account.GoTrueId,
	}

	services.ResourceService.CreateResource(&resource)

	c.Redirect(http.StatusFound, "/resources/mine")
	c.Abort()
}

func ListAllResources(c *gin.Context) {
	resources := services.ResourceService.GetAll()


	c.HTML(http.StatusOK, "resource-list-view", gin.H{
		"resources": resources,
		"user":      controllers.GetUserFromContext(c),

	})
}

func ListPrivateResources(c *gin.Context) {
	accountId := controllers.GetUserFromContext(c).Id
	resources := services.ResourceService.GetAllByAccountId(accountId)


	c.HTML(http.StatusOK, "resource-list-view-admin", gin.H{
		"resources": resources,
		"user":      controllers.GetUserFromContext(c),
	})
}

func ViewResource(c *gin.Context) {
	resourceId := c.Param("id")
	resource := services.ResourceService.GetOne(resourceId)

	c.HTML(http.StatusOK, "resource-view", gin.H{
		"resource": resource,
		"user":     controllers.GetUserFromContext(c),
	})
}

func RenderResourceQRCode(c *gin.Context) {
	resourceId := c.Param("id")

	// Generate the QRCode
	uri := fmt.Sprintf("https://www.contextual.guide/resources/%s/redirect", resourceId)
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

	c.Redirect(http.StatusFound, resource.Redirect)

	c.Abort()
}
