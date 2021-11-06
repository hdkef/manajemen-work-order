package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PUMDashboard(c *gin.Context) {
	//render pum dashboard
	c.HTML(http.StatusOK, "pum/dashboard.html", nil)
}
