package services

import (
	"manajemen-work-order/models"

	"github.com/gin-gonic/gin"
)

func SendBasicResponse(c *gin.Context, code int, ok bool, msg string) {
	c.JSON(code, models.BasicResponse{
		OK:  ok,
		Msg: msg,
	})
}
