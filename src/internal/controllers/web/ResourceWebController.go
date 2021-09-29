package webController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
	"log"
	"main/src/internal"
	"main/src/internal/controllers"
	"net/http"
	"os"
	"strconv"
	"time"
)

func RenderResourceForm(c *gin.Context) {
	c.HTML(http.StatusOK, "resource-form", gin.H{
		"user": controllers.GetUserFromContext(c),
	})
}

func CreateResource(c *gin.Context) {

	account := internal.AccountService.GetAccount(controllers.GetUserFromContext(c).Id)

	longitude, err := strconv.ParseFloat(c.PostForm("longitude"), 32)

	errList := make(internal.ValidationError)

	if err != nil {
		message := fmt.Sprintf("logitude is not in the right format : x.x")
		errList["longitude"] = append(errList["longitude"], message)

	}
	latitude, err := strconv.ParseFloat(c.PostForm("latitude"), 32)
	if err != nil {
		message := fmt.Sprintf("latitude is not in the right format : x.x")
		errList["latitude"] = append(errList["latitude"], message)

	}
	timeLayout := "2006-01-02T15:00"
	timestamp, err := time.Parse(timeLayout, c.PostForm("timestamp"))

	if err != nil {
		message := fmt.Sprintf("Timestamp is not in the right format : %s", timeLayout)
		errList["timestamp"] = append(errList["latitude"], message)
	}

	command := internal.ResourceSaveCommand{
		Title:            c.PostForm("title"),
		Description:      c.PostForm("description"),
		Timestamp:        timestamp,
		Longitude:        float32(longitude),
		Latitude:         float32(latitude),
		Redirect:         c.PostForm("Redirect"),
		CustomProperties: c.PostForm("customProperties"),
	}
	if len(errList) != 0 {
		c.HTML(http.StatusOK, "resource-form", gin.H{
			"errors": errList,
			"user":   controllers.GetUserFromContext(c),
			"model":  command,
		})
		return
	}

	resource, errorList := internal.NewResource(command, account.GoTrueId)

	fmt.Println(errorList)
	if len(*errorList) != 0 {
		c.HTML(http.StatusOK, "resource-form", gin.H{
			"errors": errorList,
			"user":   controllers.GetUserFromContext(c),
			"model":  command,
		})

		return
	}

	internal.ResourceService.CreateResource(resource)

	c.Redirect(http.StatusFound, "/resources/mine")
	c.Abort()
}

func ListAllResources(c *gin.Context) {
	resources := internal.ResourceService.GetAll()

	c.HTML(http.StatusOK, "resource-list-view", gin.H{
		"resources": resources,
		"user":      controllers.GetUserFromContext(c),
	})
}

func ListPrivateResources(c *gin.Context) {
	accountId := controllers.GetUserFromContext(c).Id
	resources := internal.ResourceService.GetAllByAccountId(accountId)

	c.HTML(http.StatusOK, "resource-list-view-admin", gin.H{
		"resources": resources,
		"user":      controllers.GetUserFromContext(c),
	})
}

func ViewResource(c *gin.Context) {
	resourceId := c.Param("id")
	resource := internal.ResourceService.GetOne(resourceId)

	c.HTML(http.StatusOK, "resource-view", gin.H{
		"resource": resource,
		"user":     controllers.GetUserFromContext(c),
	})
}

func RenderResourceQRCode(c *gin.Context) {
	resourceId := c.Param("id")

	// Generate the QRCode
	uri := fmt.Sprintf("https://%s/resources/%s/redirect", os.Getenv("APP_URL"), resourceId)

	q, err := qrcode.New(uri, qrcode.High)
	if err != nil {
		panic(err)
	}

	q.DisableBorder = true

	png, err := q.PNG(256)
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
	resource := internal.ResourceService.GetOne(resourceId)

	redirect := resource.Redirect
	if redirect == "" {
		redirect = fmt.Sprintf("%s/%s",
			os.Getenv("APP_URL"),
			resourceId)
	}

	c.Redirect(http.StatusFound, redirect)
	c.Abort()
}
