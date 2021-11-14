package views

import (
	"manajemen-work-order/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil || entity.Role != "User" {
		c.JSON(http.StatusForbidden, "forbidden")
		return
	}
	c.HTML(http.StatusOK, "user.html", nil)
}

func SuperAdmin(c *gin.Context) {
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil || entity.Role != "Super-Admin" {
		c.JSON(http.StatusForbidden, "forbidden")
		return
	}
	c.HTML(http.StatusOK, "super-admin.html", nil)
}

func BDMU(c *gin.Context) {
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil || entity.Role != "BDMU" {
		c.JSON(http.StatusForbidden, "forbidden")
		return
	}
	c.HTML(http.StatusOK, "bdmu.html", nil)
}

func BDMUP(c *gin.Context) {
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil || entity.Role != "BDMUP" {
		c.JSON(http.StatusForbidden, "forbidden")
		return
	}
	c.HTML(http.StatusOK, "bdmup.html", nil)
}

func KELA(c *gin.Context) {
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil || entity.Role != "KELA" {
		c.JSON(http.StatusForbidden, "forbidden")
		return
	}
	c.HTML(http.StatusOK, "kela.html", nil)
}

func KELB(c *gin.Context) {
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil || entity.Role != "KELB" {
		c.JSON(http.StatusForbidden, "forbidden")
		return
	}
	c.HTML(http.StatusOK, "kelb.html", nil)
}

func PPK(c *gin.Context) {
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil || entity.Role != "PPK" {
		c.JSON(http.StatusForbidden, "forbidden")
		return
	}
	c.HTML(http.StatusOK, "ppk.html", nil)
}

func PPE(c *gin.Context) {
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil || entity.Role != "PPE" {
		c.JSON(http.StatusForbidden, "forbidden")
		return
	}
	c.HTML(http.StatusOK, "ppe.html", nil)
}

func ULP(c *gin.Context) {
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil || entity.Role != "ULP" {
		c.JSON(http.StatusForbidden, "forbidden")
		return
	}
	c.HTML(http.StatusOK, "ulp.html", nil)
}
