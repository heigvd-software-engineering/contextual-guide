package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"main/src/internal/models"
	"main/src/internal/services"
	"net/http"
	"strconv"
)

func RenderTokenForm(c *gin.Context) {
	c.HTML(http.StatusOK, "token-form", gin.H{
		"user": getUserFromContext(c),
	})
}

func CreateToken(c *gin.Context) {

	account := services.AccountService.GetAccount(getUserFromContext(c).Id)


	token := models.Token{
		Name:  c.PostForm("name"),
		Value: uuid.New().String(),
		Account: *account,
		AccountId: account.GoTrueId,
	}

	services.TokenService.CreateToken(&token)

	c.HTML(http.StatusOK, "created-token-view", gin.H{
		"token": token,
		"user": getUserFromContext(c),
	})

}

func GetTokens(c *gin.Context) {
	tokens := services.TokenService.GetAll()

	c.HTML(http.StatusOK, "token-list-view", gin.H{
		"tokens": tokens,
		"user": getUserFromContext(c),
	})
}

func DeleteToken(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	services.TokenService.Delete(id)

	c.Redirect(http.StatusFound, "/tokens")
	c.Abort()
}
