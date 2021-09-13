package controllers

import (
	"github.com/gin-gonic/gin"
	"main/src/internal/services"
	"net/http"
	"strconv"
)

func GetAccount(c *gin.Context)  {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	user, apiErr := services.AccountService.GetAccount(userID)

	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK,user)
}
