package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"main/src/internal/models"
	"main/src/internal/services"
	"net/http"
	"strconv"
)

func RenderTokenForm(c *gin.Context) {
	c.HTML(http.StatusOK, "token-form", nil)
}

func CreateToken(c *gin.Context) {

	token := models.Token{
		Name:  c.PostForm("name"),
		Value: uuid.New().String(),
	}

	services.TokenService.CreateToken(&token)

	user , _ :=c.Get("user")
	c.HTML(http.StatusOK, "created-token-view", gin.H{
		"token": token,
		"user": user,
	})

}

func GetTokens(c *gin.Context) {
	fmt.Println("TOKENS")
	tokens := services.TokenService.GetAll()
	user , _ :=c.Get("user")

	c.HTML(http.StatusOK, "token-list-view", gin.H{
		"tokens": tokens,
		"user": user,
	})
}

func DeleteToken(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	services.TokenService.Delete(id)

	c.Redirect(http.StatusFound, "/tokens")
	c.Abort()
}
