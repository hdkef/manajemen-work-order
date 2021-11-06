package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PPEDashboard(c *gin.Context) {
	//render ppe dashboard
	c.HTML(http.StatusOK, "ppe/dashboard.html", nil)
}
