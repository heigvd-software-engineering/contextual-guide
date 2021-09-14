package controllers

import (
	"github.com/gin-gonic/gin"
)

func Render(c *gin.Context) {

	viewName := c.Request.RequestURI
	viewName = viewName[1:]

	if viewName == "" {
		viewName = "home"
	}

	c.HTML(200, viewName, nil)
}