package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserDashboard(c *gin.Context) {
	//render user dashboard
	c.HTML(http.StatusOK, "user/dashboard.html", nil)
}
