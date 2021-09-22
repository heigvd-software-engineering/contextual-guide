package controllers

import (
	"github.com/gin-gonic/gin"
)

type LoggedUser struct {
	Id string
	Email string
}

func GetUserFromContext(c *gin.Context) *LoggedUser {
	user , ok :=c.Get("user")
	if !ok || user == nil{
		return nil
	}

	var loggedUser LoggedUser
	loggedUser = user.(LoggedUser)

	return &loggedUser
}

func Render(c *gin.Context) {

	viewName := c.Request.RequestURI
	viewName = viewName[1:]

	if viewName == "" {
		viewName = "home"
	}
	c.HTML(200, viewName, gin.H{
		"user": GetUserFromContext(c),
	})
}

func RenderErrorPage(code int, message string, c *gin.Context){
	c.HTML(code,"error",gin.H{
		"code":    code,
		"message": message,
		"user":    GetUserFromContext(c),
	})
	c.Abort()
}