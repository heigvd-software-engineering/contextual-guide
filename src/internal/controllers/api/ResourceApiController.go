package apiController

import (
	"github.com/gin-gonic/gin"
	"main/src/internal/controllers"
	"main/src/internal/models"
	"main/src/internal/services"
	"net/http"
)



// swagger:route POST /resource Resource resourceSaveCommand
// Create a new Resource
// responses:
//   201:
//     description: Resource successfully created
//   401:
//     description: Unauthorized
//   422: validationError
func CreateResource(c *gin.Context) {

	account := services.AccountService.GetAccount(controllers.GetUserFromContext(c).Id)

	var command models.ResourceSaveCommand

	if err := c.ShouldBindJSON(&command); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	resource, errorList := models.NewResource(command, account.GoTrueId)

	if errorList != nil {
		c.JSON(http.StatusBadRequest, models.ErrorDTO{Errors: errorList})
		return
	}
	services.ResourceService.CreateResource(resource)

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
