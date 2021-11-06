package controllers

import (
	"manajemen-work-order/models"
	"manajemen-work-order/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	if c.Request.Method == http.MethodPost {
		handleLoginPost(c)
		return
	}
	//if there are Authorization cookies do redirect
	val, exist := c.Get("User")
	if exist {
		user := val.(models.User)
		redirectByRole(c, user.Role)
	}
	//render login page
	c.HTML(http.StatusOK, "other/login.html", nil)
}

func redirectByRole(c *gin.Context, role string) {
	switch role {
	case "User":
		c.Redirect(http.StatusTemporaryRedirect, "/user-dashboard")
	}
}

func handleLoginPost(c *gin.Context) {
	user := models.User{}
	err := user.Authenticate(c)
	if err != nil {
		utils.Response(c, http.StatusUnauthorized, false, err.Error())
		return
	}
	//TOBE
	switch user.Role {
	case "User":
		utils.Response(c, http.StatusOK, true, "http://localhost:8080/user-dashboard")
	}
}
