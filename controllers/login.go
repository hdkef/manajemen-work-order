package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	//render login page
	c.HTML(http.StatusOK, "other/login.html", nil)
}
