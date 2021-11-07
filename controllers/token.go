package controllers

import (
	"manajemen-work-order/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Token(c *gin.Context) {
	cookieString, err := c.Cookie("Authorization")
	if err != nil {
		utils.Response(c, http.StatusUnauthorized, false, err.Error())
		return
	}

	utils.Response(c, http.StatusOK, true, cookieString)
}