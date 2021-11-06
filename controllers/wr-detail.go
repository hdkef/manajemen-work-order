package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WRDetail(c *gin.Context) {
	//render wr-detail dashboard
	c.HTML(http.StatusOK, "other/wr-detail.html", nil)
}
