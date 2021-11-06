package controllers

import (
	"manajemen-work-order/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserDashboard(c *gin.Context) {
	//if not authenticated then dont render pages
	val, exist := c.Get("User")
	if !exist {
		return
	}

	user := val.(models.User)

	//role must be user
	if user.Role != "User" {
		c.JSON(http.StatusUnauthorized, "role bukan User")
		return
	}

	//render user dashboard
	c.HTML(http.StatusOK, "user/dashboard.html", nil)
}
