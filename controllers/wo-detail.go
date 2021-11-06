package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WODetail(c *gin.Context) {
	//render wo-detail dashboard
	c.HTML(http.StatusOK, "other/wo-detail.html", nil)
}
