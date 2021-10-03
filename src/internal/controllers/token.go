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
	account := models.GetOrCreateAccount(GetUserFromContext(c).Id)

	println(account.GoTrueId)

	value := models.CreateTokenValue()
	token := models.Token{
		Name:      c.PostForm("name"),
		Hash:      models.HashTokenValue(value),
		Account:   *account,
		AccountId: account.GoTrueId,
	}

	models.CreateToken(&token)

	c.HTML(http.StatusOK, "created-token-view", gin.H{
		"name": token.Name,
		"value": value,
		"user":  GetUserFromContext(c),
	})
}

func GetTokens(c *gin.Context) {
	accountId := GetUserFromContext(c).Id

	tokens := models.ListTokenByAccountId(accountId)

	c.HTML(http.StatusOK, "token-list-view-admin", gin.H{
		"tokens": tokens,
		"user":   GetUserFromContext(c),
	})
}

func DeleteToken(c *gin.Context) {
	hash := c.Param("hash")
	models.DeleteToken(hash)

	c.Redirect(http.StatusFound, "/tokens")
	c.Abort()
}
