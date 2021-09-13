package controllers

import (
	"github.com/gin-gonic/gin"
	"main/src/internal/services"
	"net/http"
	"strconv"
)

func GetAccount(c *gin.Context)  {
	id, _ := strconv.ParseInt(c.Param("accountId"), 10, 64)

	user := services.AccountService.GetAccount(id)

	c.JSON(http.StatusOK, user)
}

func RenderForm(c *gin.Context) {
	c.HTML(http.StatusOK,"account-form.html",gin.H{
		"title": "patate",
	})
}
