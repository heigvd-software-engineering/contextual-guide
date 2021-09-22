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
	c.HTML(http.StatusOK, "resource-form", gin.H{
		"user": getUserFromContext(c),
	})
}

func CreateResource(c *gin.Context) {

	account := services.AccountService.GetAccount(getUserFromContext(c).Id)


	resource := models.Resource{
		Uuid: shortuuid.New(),
		Content: c.PostForm("resource"),
		Account: *account,
		AccountId: account.GoTrueId,
	}

	services.ResourceService.CreateResource(&resource)

	c.Redirect(http.StatusFound, "/resources/mine")
	c.Abort()
}

func ListAllResources(c *gin.Context) {
	resources := services.ResourceService.GetAll()


	c.HTML(http.StatusOK, "public-resource-list-view", gin.H{
		"resources": resources,
		"user": getUserFromContext(c),

	})
}

func ListPrivateResources(c *gin.Context) {
	accountId := getUserFromContext(c).Id
	resources := services.ResourceService.GetAllByAccountId(accountId)


	c.HTML(http.StatusOK, "private-resource-list-view", gin.H{
		"resources": resources,
		"user": getUserFromContext(c),
	})
}

func ViewResource(c *gin.Context) {
	resourceId := c.Param("id")
	resource := services.ResourceService.GetOne(resourceId)

	c.HTML(http.StatusOK, "resource-view", gin.H{
		"resource": resource,
		"user": getUserFromContext(c),
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
		c.Redirect(http.StatusFound, content.Redirect)
	} else {
		uri := fmt.Sprintf("/resources/%s", resourceId)
		c.Redirect(http.StatusFound, uri)
	}

	c.Abort()
}

type ResourceSaveCommand struct {
	Document    string `json:"document"`
}

type ResourceDTO struct {
	Uuid		string `json:"uuid"`
	Document    string `json:"document"`
}
// swagger:route POST /resource Resource resourceSaveCommand
// Create a new Resource
// responses:
//   201:
//     description: Resource successfully created
//   401:
//     description: Unauthorized
//   422:
//     description: The model validation failed
func CreateResourceApi(c *gin.Context) {

	account := services.AccountService.GetAccount(getUserFromContext(c).Id)

	var Command ResourceSaveCommand

	if err := c.ShouldBindJSON(&Command); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}


	resource := models.Resource{
		Uuid: shortuuid.New(),
		Content: Command.Document,
		Account: *account,
		AccountId: account.GoTrueId,
	}

	services.ResourceService.CreateResource(&resource)

	c.JSON(http.StatusCreated,nil)
}

// swagger:route GET /resource Resource Resource
// Get all resources scoped by the apikey
// responses:
//   200: resourceDTOList
//   401:
//     description: Unauthorized
func ListPrivateResourcesApi(c *gin.Context) {
	accountId := getUserFromContext(c).Id
	resources := services.ResourceService.GetAllByAccountId(accountId)

	c.JSON(http.StatusOK, resources)
}
// swagger:route GET /resource/:uuid Resource resourceGetById
// Get one resource by id
// responses:
//   200: resourceDTO
//   401:
//     description: Unauthorized
func ViewResourceApi(c *gin.Context) {
	resourceId := c.Param("id")
	resource := services.ResourceService.GetOne(resourceId)

	c.JSON(http.StatusOK,resource)
}
