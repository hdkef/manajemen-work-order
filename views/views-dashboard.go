package views

import (
	"manajemen-work-order/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil || entity.Role != "USER" {
		c.JSON(http.StatusForbidden, "forbidden")
		return
	}
	c.HTML(http.StatusOK, "user.html", entity.Fullname)
}

func SuperAdmin(c *gin.Context) {
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil || entity.Role != "Super-Admin" {
		c.JSON(http.StatusForbidden, "forbidden")
		return
	}
	c.HTML(http.StatusOK, "super-admin.html", entity.Fullname)
}

func BDMU(c *gin.Context) {
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil || entity.Role != "BDMU" {
		c.JSON(http.StatusForbidden, "forbidden")
		return
	}
	c.HTML(http.StatusOK, "bdmu.html", entity.Fullname)
}

func BDMUP(c *gin.Context) {
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil || entity.Role != "BDMUP" {
		c.JSON(http.StatusForbidden, "forbidden")
		return
	}
	c.HTML(http.StatusOK, "bdmup.html", entity.Fullname)
}

func KELA(c *gin.Context) {
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil || entity.Role != "KELA" {
		c.JSON(http.StatusForbidden, "forbidden")
		return
	}
	c.HTML(http.StatusOK, "kela.html", entity.Fullname)
}

func KELB(c *gin.Context) {
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil || entity.Role != "KELB" {
		c.JSON(http.StatusForbidden, "forbidden")
		return
	}
	c.HTML(http.StatusOK, "kelb.html", entity.Fullname)
}

func PPK(c *gin.Context) {
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil || entity.Role != "PPK" {
		c.JSON(http.StatusForbidden, "forbidden")
		return
	}
	c.HTML(http.StatusOK, "ppk.html", entity.Fullname)
}

func PPE(c *gin.Context) {
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil || entity.Role != "PPE" {
		c.JSON(http.StatusForbidden, "forbidden")
		return
	}
	c.HTML(http.StatusOK, "ppe.html", entity.Fullname)
}

func ULP(c *gin.Context) {
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil || entity.Role != "ULP" {
		c.JSON(http.StatusForbidden, "forbidden")
		return
	}
	c.HTML(http.StatusOK, "ulp.html", entity.Fullname)
}
