package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WOProgress(c *gin.Context) {
	//render wo-progress dashboard
	c.HTML(http.StatusOK, "other/wo-progress.html", nil)
}
