package controllers

import (
	"github.com/gin-gonic/gin"
)

type LoggedUser struct {
	Id string
	Email string
}

func Render(c *gin.Context) {

	viewName := c.Request.RequestURI
	viewName = viewName[1:]

	if viewName == "" {
		viewName = "home"
	}
	var loggedUser *LoggedUser
	user , ok :=c.Get("user")

	if !ok {
		loggedUser = user.(*LoggedUser)
	}

	c.HTML(200, viewName, gin.H{
		"user": loggedUser,
	})
}

func RenderErrorPage(code int, message string, c *gin.Context){
	user , _ :=c.Get("user")
	c.HTML(code,"error",gin.H{
		"code": code,
		"message": message,
		"user": user,
	})
	c.Abort()
}