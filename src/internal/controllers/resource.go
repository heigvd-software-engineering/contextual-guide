package controllers

import (
	err "errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	qrcode "github.com/skip2/go-qrcode"
	"log"
	"main/src/internal/models"
	"strconv"

	"net/http"
	"os"
)

func LatestResources(c *gin.Context) {
	resources := models.LatestResources()
	c.HTML(http.StatusOK, "latest", gin.H{
		"resources": resources,
		"user":      GetUserFromContext(c),
	})
}

func ListResources(c *gin.Context) {
	user := GetUserFromContext(c)
	resources := models.GetAllResourceByAccountId(user.Id)
	c.HTML(http.StatusOK, "resource-list", gin.H{
		"resources": resources,
		"user":      user,
	})
}

func GetResource(c *gin.Context) {
	user := GetUserFromContext(c)

	uuid := c.Param("uuid")
	resource := models.ReadResource(uuid)

	c.HTML(http.StatusOK, "resource-view", gin.H{
		"resource": resource,
		"user":     user,
	})
}

func CreateResourceForm(c *gin.Context) {
	user := GetUserFromContext(c)

	title := "New signal"
	action := "/resources"
	errors := make(map[string]string)
	resource := &models.Resource{}

	c.HTML(http.StatusOK, "resource-form", gin.H{
		"title":    title,
		"action":   action,
		"user":     user,
		"resource": resource,
		"errors":   errors,
	})
}

func CreateResource(c *gin.Context) {
	user := GetUserFromContext(c)

	var resource models.Resource
	error := c.ShouldBind(&resource)
	fmt.Println(resource.Referenced)


	errors := make(map[string]string)
	if error != nil {
		var validationErrors validator.ValidationErrors
		if err.As(error, &validationErrors) {
			for _, b := range validationErrors {
				errors[b.Field()] = b.Error()
			}
		}
		var title = "New signal"
		var action = "/resources"
		c.HTML(http.StatusOK, "resource-form", gin.H{
			"title":    title,
			"action":   action,
			"user":     user,
			"resource": resource,
			"errors":   errors,
		})
		return
	}

	models.CreateResource(user.Id, &resource)

	c.Redirect(http.StatusFound, "/resources")
	c.Abort()
}

func UpdateResourceForm(c *gin.Context) {
	user := GetUserFromContext(c)
	uuid := c.Param("uuid")

	title := "Edit signal"
	action := "/resources/" + uuid
	resource := models.ReadResource(uuid)
	errors := make(map[string]string)

	c.HTML(http.StatusOK, "resource-form", gin.H{
		"title":    title,
		"action":   action,
		"user":     user,
		"resource": resource,
		"errors":   errors,
	})
}

func UpdateResource(c *gin.Context) {
	user := GetUserFromContext(c)

	resource := models.Resource{}
	error := c.ShouldBind(&resource)
	if error != nil {
		fmt.Println(error)
	}

	resource.Uuid = c.Param("uuid")
	resource.AccountId = user.Id

	models.UpdateResource(user.Id, &resource)

	c.HTML(http.StatusOK, "resource-view", gin.H{
		"resource": resource,
		"user":     user,
	})
}

func DeleteResource(c *gin.Context) {
	user := GetUserFromContext(c)
	uuid := c.Param("uuid")

	models.DeleteResource(user.Id, uuid)

	c.Redirect(http.StatusFound, "/resources")
	c.Abort()
}

func RenderResourceQRCode(c *gin.Context) {
	uuid := c.Param("uuid")

	// Check the size of the QRCode
	size, err := strconv.Atoi(c.Param("size"))
	if err != nil || !(size >= 128 && size <= 512) {
		c.AbortWithStatus(404)
		return
	}

	// Generate the QRCode
	uri := fmt.Sprintf("https://%s/resources/%s/redirect", os.Getenv("APP_URL"), uuid)

	q, err := qrcode.New(uri, qrcode.High)
	if err != nil {
		panic(err)
	}

	q.DisableBorder = true

	png, err := q.PNG(size)
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
	resourceId := c.Param("uuid")
	resource := models.ReadResource(resourceId)

	var redirect = fmt.Sprintf("%s/%s",
		os.Getenv("APP_URL"),
		resourceId)

	if len(resource.Redirect) > 0 {
		redirect = resource.Redirect
	}

	c.Redirect(http.StatusFound, redirect)
	c.Abort()
}
