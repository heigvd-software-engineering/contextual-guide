package apiController

import (
	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v3"
	"main/src/internal/controllers"
	"main/src/internal/models"
	"main/src/internal/services"
	"net/http"
	"time"
)

type Content struct {
	Redirect string
}


type ResourceSaveCommand struct {
	Title    string `json:"title"`
	Description    string `json:"description"`
	Timestamp    time.Time `json:"timestamp"`
	Longitude    float32 `json:"longitude"`
	Latitude    float32 `json:"latitude"`
	Redirect    string `json:"redirect"`
	CustomProperties string `json:"CustomProperties"`
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
func CreateResource(c *gin.Context) {

	account := services.AccountService.GetAccount(controllers.GetUserFromContext(c).Id)

	var command ResourceSaveCommand

	if err := c.ShouldBindJSON(&command); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}


	resource := models.Resource{
		Uuid: shortuuid.New(),
		Title: command.Title,
		Description: command.Description,
		Timestamp: command.Timestamp,
		Longitude: command.Longitude,
		Latitude: command.Latitude,
		Redirect: command.Redirect,
		CustomProperties: command.CustomProperties,
		AccountId: account.GoTrueId,
	}

	services.ResourceService.CreateResource(&resource)

	c.JSON(http.StatusCreated,nil)
}

// swagger:route GET /resource Resource Resource
// Get all resources scoped by the apikey
// responses:
//   200: resourceList
//   401:
//     description: Unauthorized
func ListPrivateResources(c *gin.Context) {
	accountId := controllers.GetUserFromContext(c).Id
	resources := services.ResourceService.GetAllByAccountId(accountId)

	c.JSON(http.StatusOK, resources)
}
// swagger:route GET /resource/:uuid Resource resourceGetById
// Get one resource by id
// responses:
//   200: resource
//   401:
//     description: Unauthorized
func ViewResource(c *gin.Context) {
	resourceId := c.Param("id")
	resource := services.ResourceService.GetOne(resourceId)

	c.JSON(http.StatusOK,resource)
}
