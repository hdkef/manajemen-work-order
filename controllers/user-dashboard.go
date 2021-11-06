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
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}

	user := val.(models.User)

	//role must be user
	if user.Role != "User" {
		//remove cookie
		c.SetCookie(models.AUTH_COOKIE_NAME, "", -1, "/", "", false, false)
		//response
		c.JSON(http.StatusUnauthorized, "role bukan User")
		return
	}

	//render user dashboard
	c.HTML(http.StatusOK, "user/dashboard.html", nil)
}
