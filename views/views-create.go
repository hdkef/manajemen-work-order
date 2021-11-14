package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEntity(c *gin.Context) {
	c.HTML(http.StatusOK, "create-entity.html", nil)
}

func CreatePPP(c *gin.Context) {
	c.HTML(http.StatusOK, "create-ppp.html", nil)
}

func CreateRP(c *gin.Context) {
	c.HTML(http.StatusOK, "create-rp.html", nil)
}

func CreatePerkiraanBiaya(c *gin.Context) {
	c.HTML(http.StatusOK, "create-perkiraan-biaya.html", nil)
}

func CreatePengadaan(c *gin.Context) {
	c.HTML(http.StatusOK, "create-pengadaan.html", nil)
}

func CreateSPK(c *gin.Context) {
	c.HTML(http.StatusOK, "create-spk.html", nil)
}
