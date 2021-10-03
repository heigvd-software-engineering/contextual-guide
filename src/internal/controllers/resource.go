package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v3"
	qrcode "github.com/skip2/go-qrcode"
	"log"
	"main/src/internal/models"
	"strconv"

	"net/http"
	"os"
)

func ResourceForm(c *gin.Context) {
	var title string = "New event"
	var action string = "/resources"
	var resource models.Resource

	uuid := c.Param("uuid")
	if uuid != "" {
		title = "Edit event"
		action = "/resources/" + uuid
		resource = *models.ReadResource(uuid)
	}

	c.HTML(http.StatusOK, "resource-form", gin.H{
		"title":    title,
		"action":   action,
		"resource": resource,
		"user":     GetUserFromContext(c),
	})
}

func CreateResource(c *gin.Context) {
	user := GetUserFromContext(c)
	resource := models.Resource{
		Uuid:      shortuuid.New(),
		AccountId: user.Id,
	}

	error := c.ShouldBind(&resource)
	if error != nil {
		fmt.Println(error)
	}

	models.CreateResource(&resource)

	c.Redirect(http.StatusFound, "/resources")
	c.Abort()
}

func ReadResource(c *gin.Context) {
	resourceId := c.Param("uuid")

	resource := models.ReadResource(resourceId)

	c.HTML(http.StatusOK, "resource-view", gin.H{
		"resource": resource,
		"user":     GetUserFromContext(c),
	})
}

func UpdateResource(c *gin.Context) {
	uuid := c.Param("uuid")
	user := GetUserFromContext(c)
	resource := models.Resource{
		Uuid:      uuid,
		AccountId: user.Id,
	}

	error := c.ShouldBind(&resource)
	if error != nil {
		fmt.Println(error)
	}

	models.UpdateResource(&resource)

	c.HTML(http.StatusOK, "resource-view", gin.H{
		"resource": resource,
		"user":     GetUserFromContext(c),
	})
}

func DeleteResource(c *gin.Context) {
	uuid := c.Param("uuid")

	models.DeleteResource(uuid)

	c.Redirect(http.StatusFound, "/resources")
	c.Abort()
}

func Registry(c *gin.Context) {
	resources := models.GetAllResources()
	c.HTML(http.StatusOK, "resource-list-view", gin.H{
		"resources": resources,
		"user":      GetUserFromContext(c),
	})
}

func ListResources(c *gin.Context) {
	accountId := GetUserFromContext(c).Id
	resources := models.GetAllResourceByAccountId(accountId)

	c.HTML(http.StatusOK, "resource-list-view-admin", gin.H{
		"resources": resources,
		"user":      GetUserFromContext(c),
	})
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

	redirect := resource.Redirect
	if redirect == "" {
		redirect = fmt.Sprintf("%s/%s",
			os.Getenv("APP_URL"),
			resourceId)
	}

	c.Redirect(http.StatusFound, redirect)
	c.Abort()
}
