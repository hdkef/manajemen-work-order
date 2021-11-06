package utils

import "github.com/gin-gonic/gin"

func Response(c *gin.Context, code int, ok bool, msg string) {
	payload := struct {
		OK  bool   `json:"ok"`
		Msg string `json:"msg"`
	}{
		OK:  ok,
		Msg: msg,
	}
	c.JSON(code, payload)
}
