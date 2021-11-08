package controllers

import (
	"manajemen-work-order/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PPEDashboard(c *gin.Context) {
	//if not authenticated then dont render pages
	val, exist := c.Get("User")
	if !exist {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}

	user := val.(models.User)

	//role must be PUM
	if user.Role != "PPE" {
		//remove cookie
		c.SetCookie(models.AUTH_COOKIE_NAME, "", -1, "/", "", false, false)
		//response
		c.JSON(http.StatusUnauthorized, "role bukan PPE")
		return
	}
	//render ppe dashboard
	c.HTML(http.StatusOK, "ppe/dashboard.html", nil)
}
