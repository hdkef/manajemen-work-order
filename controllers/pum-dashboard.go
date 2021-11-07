package controllers

import (
	"manajemen-work-order/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PUMDashboard(c *gin.Context) {
	//if not authenticated then dont render pages
	val, exist := c.Get("User")
	if !exist {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}

	user := val.(models.User)

	//role must be PUM
	if user.Role != "PUM" {
		//remove cookie
		c.SetCookie(models.AUTH_COOKIE_NAME, "", -1, "/", "", false, false)
		//response
		c.JSON(http.StatusUnauthorized, "role bukan PUM")
		return
	}
	//render pum dashboard
	c.HTML(http.StatusOK, "pum/dashboard.html", nil)
}
