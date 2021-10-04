package controllers

import (
	"github.com/gin-gonic/gin"
	"main/src/internal/models"
	"net/http"
)

func RenderTokenForm(c *gin.Context) {
	c.HTML(http.StatusOK, "token-form", gin.H{
		"user": GetUserFromContext(c),
	})
}

func CreateToken(c *gin.Context) {
	account := GetUserFromContext(c)

	value := models.CreateTokenValue()
	token := &models.Token{
		Name: c.PostForm("name"),
		Hash: models.HashTokenValue(value),
	}

	models.CreateToken(account.Id, token)

	c.HTML(http.StatusOK, "token-view", gin.H{
		"name":  token.Name,
		"value": value,
		"user":  GetUserFromContext(c),
	})
}

func GetTokens(c *gin.Context) {
	account := GetUserFromContext(c)

	tokens := models.ListTokenByAccountId(account.Id)

	c.HTML(http.StatusOK, "token-list", gin.H{
		"tokens": tokens,
		"user":   GetUserFromContext(c),
	})
}

func DeleteToken(c *gin.Context) {
	account := GetUserFromContext(c)
	hash := c.Param("hash")
	models.DeleteToken(account.Id, hash)

	c.Redirect(http.StatusFound, "/tokens")
	c.Abort()
}
