package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PPKDashboard(c *gin.Context) {
	//render ppk dashboard
	c.HTML(http.StatusOK, "ppk/dashboard.html", nil)
}
