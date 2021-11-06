package middlewares

import (
	"manajemen-work-order/models"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	user := models.User{}
	err := user.ValidateTokenFromCookies(c)
	if err != nil {
		c.Next()
		return
	}
	c.Set("User", user)
	c.Next()
}
