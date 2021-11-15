package views

import (
	"manajemen-work-order/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEntity(c *gin.Context) {
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		c.JSON(http.StatusForbidden, "forbidden")
		return
	}
	if entity.Role != "Super-Admin" {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	c.HTML(http.StatusOK, "create-entity.html", nil)
}

func CreatePPP(c *gin.Context) {
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		c.JSON(http.StatusForbidden, "forbidden")
		return
	}
	if entity.Role != "USER" {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	c.HTML(http.StatusOK, "create-ppp.html", nil)
}

func CreateRP(c *gin.Context) {
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		c.JSON(http.StatusForbidden, "forbidden")
		return
	}
	if entity.Role != "KELB" {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	c.HTML(http.StatusOK, "create-rp.html", struct {
		ID    string
		PPPID string
	}{
		ID:    c.Params.ByName("id"),
		PPPID: c.Params.ByName("ppp_id"),
	})
}

func CreatePerkiraanBiaya(c *gin.Context) {
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		c.JSON(http.StatusForbidden, "forbidden")
		return
	}
	if entity.Role != "PPK" {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	c.HTML(http.StatusOK, "create-perkiraan-biaya.html", struct {
		ID   string
		RPID string
	}{
		ID:   c.Params.ByName("id"),
		RPID: c.Params.ByName("rp_id"),
	})
}

func CreatePengadaan(c *gin.Context) {
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		c.JSON(http.StatusForbidden, "forbidden")
		return
	}
	x := entity.Role != "ULP"
	if x || entity.Role != "PPE" {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	c.HTML(http.StatusOK, "create-pengadaan.html", struct {
		ID               string
		PerkiraanBiayaID string
	}{
		ID:               c.Params.ByName("id"),
		PerkiraanBiayaID: c.Params.ByName("perkiraan_biaya_id"),
	})
}

func CreateSPK(c *gin.Context) {
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		c.JSON(http.StatusForbidden, "forbidden")
		return
	}
	if entity.Role != "PPK" {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	c.HTML(http.StatusOK, "create-spk.html", struct {
		ID          string
		PengadaanID string
	}{
		ID:          c.Params.ByName("id"),
		PengadaanID: c.Params.ByName("pengadaan_id"),
	})
}
