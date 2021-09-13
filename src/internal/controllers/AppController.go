package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Render(c *gin.Context) {

	viewName := c.Request.RequestURI

	fmt.Println(viewName)

	viewName = viewName[1:]

	fmt.Println(viewName)

	if viewName == "" {
		viewName = "home"
	}

	c.HTML(200, viewName, gin.H{
		"title": "Html5 Template Engine",
	})
}