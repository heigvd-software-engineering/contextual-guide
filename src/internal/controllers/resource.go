package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
	"log"
	"main/src/internal/models"
	"main/src/internal/services"
	"net/http"
	"os"
	"strconv"
	"time"
)

func RenderResourceForm(c *gin.Context) {
	c.HTML(http.StatusOK, "resource-form", gin.H{
		"user": GetUserFromContext(c),
	})
}

func CreateResource(c *gin.Context) {

	account := services.GetAccount(GetUserFromContext(c).Id)

	longitude, err := strconv.ParseFloat(c.PostForm("longitude"), 32)

	errList := make(models.ValidationError)

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

	command := ResourceSaveCommand{
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
			"user":   GetUserFromContext(c),
			"model":  command,
		})
		return
	}

	resource, errorList := NewResource(command, account.GoTrueId)

	fmt.Println(errorList)
	if len(*errorList) != 0 {
		c.HTML(http.StatusOK, "resource-form", gin.H{
			"errors": errorList,
			"user":   GetUserFromContext(c),
			"model":  command,
		})

		return
	}

	services.CreateResource(resource)

	c.Redirect(http.StatusFound, "/resources")
	c.Abort()
}

func Registry(c *gin.Context) {
	resources := services.GetAllResource()
	c.HTML(http.StatusOK, "resource-list-view", gin.H{
		"resources": resources,
		"user":      GetUserFromContext(c),
	})
}

func ListResources(c *gin.Context) {
	accountId := GetUserFromContext(c).Id
	resources := services.GetAllResourceByAccountId(accountId)

	c.HTML(http.StatusOK, "resource-list-view-admin", gin.H{
		"resources": resources,
		"user":      GetUserFromContext(c),
	})
}

func ViewResource(c *gin.Context) {
	resourceId := c.Param("id")
	resource := services.GetResource(resourceId)

	c.HTML(http.StatusOK, "resource-view", gin.H{
		"resource": resource,
		"user":     GetUserFromContext(c),
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
	resource := services.GetResource(resourceId)

	redirect := resource.Redirect
	if redirect == "" {
		redirect = fmt.Sprintf("%s/%s",
			os.Getenv("APP_URL"),
			resourceId)
	}

	c.Redirect(http.StatusFound, redirect)
	c.Abort()
}
