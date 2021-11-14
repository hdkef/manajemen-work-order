package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PPPDetail(c *gin.Context) {
	c.HTML(http.StatusOK, "ppp-detail.html", nil)
}
