package controllers

import (
	"github.com/gin-gonic/gin"
	"main/src/internal/models"
	"net/http"
)

// swagger:route GET /resource Resource getResourceList
// Gets all resources scoped by the apikey
// responses:
//   200: resourceList
//   401:
//     description: Unauthorized
func ListResourceApi(c *gin.Context) {
	accountId := GetUserFromContext(c).Id
	resources := models.GetAllResourceByAccountId(accountId)

	c.JSON(http.StatusOK, resources)
}

// swagger:route POST /resource Resource createResource
// Creates a new Resource
// responses:
//   200: resource
//   401:
//     description: Unauthorized
//   422: validationError
func CreateResourceApi(c *gin.Context) {
	account := GetUserFromContext(c)

	var resource models.Resource
	if err := c.ShouldBindJSON(&resource); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	resource.AccountId = account.Id

	models.CreateResource(account.Id, &resource)

	c.JSON(http.StatusOK, resource)
}

// swagger:route PUT /resource/:uuid Resource updateResource
// Updates an existing Resource
// responses:
//   200: resource
//   401:
//     description: Unauthorized
//   422: validationError
func UpdateResourceApi(c *gin.Context) {
	account := GetUserFromContext(c)
	uuid := c.Param("uuid")

	resource := models.ReadResource(uuid)
	error := c.ShouldBindJSON(&resource)
	if error != nil {
		c.JSON(http.StatusBadRequest, error.Error())
		return
	}

	resource.Uuid = c.Param("uuid")
	resource.AccountId = account.Id

	models.UpdateResource(account.Id, resource)

	c.JSON(http.StatusOK, resource)
}

// swagger:route GET /resource/:uuid Resource getResource
// Gets resource by uuid
// responses:
//   200: resource
//   401:
//     description: Unauthorized
func GetResourceApi(c *gin.Context) {
	uuid := c.Param("uuid")
	resource := models.ReadResource(uuid)

	c.JSON(http.StatusOK, resource)
}

// swagger:route DELETE /resource/:uuid Resource deleteResource
// Deletes resource by uuid
// responses:
//   200: resource
//   401:
//     description: Unauthorized
func DeleteResourceApi(c *gin.Context) {
	uuid := c.Param("uuid")
	resource := models.ReadResource(uuid)

	c.JSON(http.StatusOK, resource)
}

