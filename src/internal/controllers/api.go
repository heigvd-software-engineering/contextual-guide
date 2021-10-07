package controllers

import (
	"github.com/gin-gonic/gin"
	"main/src/internal/models"
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
func PostResource(c *gin.Context) {
	account := models.GetOrCreateAccount(GetUserFromContext(c).Id)

	var resource models.Resource
	if err := c.ShouldBindJSON(&resource); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	models.CreateResource(account.GoTrueId, &resource)

	c.JSON(http.StatusCreated, nil)
}

// swagger:route GET /resource Resource Resource
// Get all resources scoped by the apikey
// responses:
//   200: resourceList
//   401:
//     description: Unauthorized
func GetResources(c *gin.Context) {
	accountId := GetUserFromContext(c).Id
	resources := models.GetAllResourceByAccountId(accountId)

	c.JSON(http.StatusOK, resources)
}

// swagger:route GET /resource/:uuid Resource resourceGetById
// Get one resource by id
// responses:
//   200: resource
//   401:
//     description: Unauthorized
func GetResource(c *gin.Context) {
	resourceId := c.Param("id")
	resource := models.ReadResource(resourceId)

	c.JSON(http.StatusOK, resource)
}

