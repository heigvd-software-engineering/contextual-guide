package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v3"
	"main/src/internal/models"
	"main/src/internal/services"
	"net/http"
	"strconv"
)

func RenderTokenForm(c *gin.Context) {
	c.HTML(http.StatusOK, "token-form", gin.H{
		"user": GetUserFromContext(c),
	})
}

func CreateToken(c *gin.Context) {
	account := services.GetAccount(GetUserFromContext(c).Id)

	// the hash of the token
	token := models.Token{
		Name:  c.PostForm("name"),
		Hash: shortuuid.New(),
		Account: *account,
		AccountId: account.GoTrueId,
	}

	services.CreateToken(&token)

	c.HTML(http.StatusOK, "created-token-view", gin.H{
		"token": token,
		"user":  GetUserFromContext(c),
	})

}

func GetTokens(c *gin.Context) {
	accountId := GetUserFromContext(c).Id

	tokens := services.ListTokenByAccountId(accountId)

	c.HTML(http.StatusOK, "token-list-view-admin", gin.H{
		"tokens": tokens,
		"user":   GetUserFromContext(c),
	})
}

func DeleteToken(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	services.DeleteToken(id)

	c.Redirect(http.StatusFound, "/tokens")
	c.Abort()
}
