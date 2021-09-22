package webController

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v3"
	qrcode "github.com/skip2/go-qrcode"
	"log"
	"main/src/internal/controllers"
	apiController "main/src/internal/controllers/api"
	"main/src/internal/models"
	"main/src/internal/services"
	"net/http"
)

type Content struct {
	Redirect string
}

func RenderResourceForm(c *gin.Context) {
	c.HTML(http.StatusOK, "resource-form", gin.H{
		"user": controllers.GetUserFromContext(c),
	})
}

func CreateResource(c *gin.Context) {

	account := services.AccountService.GetAccount(controllers.GetUserFromContext(c).Id)
	
	command := apiController.ResourceSaveCommand{
		Title:       c.PostForm("title"),
		Description: c.PostForm("description"),
		Timestamp:   c.PostForm("timestamp"),
		Longitude:   c.PostForm("longitude"),
		Latitude:    c.PostForm("latitude"),
		Redirect:    c.PostForm("redirect"),
	}

	document, _ := json.Marshal(command)

	resource := models.Resource{
		Uuid: shortuuid.New(),
		Document: string(document),
		AccountId: account.GoTrueId,
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
	blob := []byte(resource.Document)
	var content Content
	err := json.Unmarshal(blob, &resource)
	if err != nil {
		log.Println(err)
		c.String(500, "Unable to unmarshal resource")
		return
	}

	if content.Redirect != "" {
		c.Redirect(http.StatusFound, content.Redirect)
	} else {
		uri := fmt.Sprintf("/resources/%s", resourceId)
		c.Redirect(http.StatusFound, uri)
	}

	c.Abort()
}
