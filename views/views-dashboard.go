package views

import (
	"manajemen-work-order/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
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
	c.HTML(http.StatusOK, "bdmu.html", nil)
}

func BDMUP(c *gin.Context) {
	c.HTML(http.StatusOK, "bdmup.html", nil)
}

func KELA(c *gin.Context) {
	c.HTML(http.StatusOK, "kela.html", nil)
}

func KELB(c *gin.Context) {
	c.HTML(http.StatusOK, "kelb.html", nil)
}

func PPK(c *gin.Context) {
	c.HTML(http.StatusOK, "ppk.html", nil)
}

func PPE(c *gin.Context) {
	c.HTML(http.StatusOK, "ppe.html", nil)
}

func ULP(c *gin.Context) {
	c.HTML(http.StatusOK, "ulp.html", nil)
}
