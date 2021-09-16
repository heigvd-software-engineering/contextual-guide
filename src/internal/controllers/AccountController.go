package controllers

import (
	"github.com/gin-gonic/gin"
	"main/src/internal/services"
	"net/http"
)


func GetAccount(c *gin.Context)  {
	id := c.PostForm("accountId")
	user := services.AccountService.GetAccount(id)
	c.JSON(http.StatusOK, user)
}
