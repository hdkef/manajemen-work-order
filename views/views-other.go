package views

import (
	"manajemen-work-order/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		c.HTML(http.StatusOK, "login.html", nil)
		return
	}
	//send response based on role
	switch entity.Role {
	case "USER":
		c.Redirect(http.StatusTemporaryRedirect, "/user")
	case "BDMU":
		c.Redirect(http.StatusTemporaryRedirect, "/bdmu")
	case "BDMUP":
		c.Redirect(http.StatusTemporaryRedirect, "/bdmup")
	case "KELA":
		c.Redirect(http.StatusTemporaryRedirect, "/kela")
	case "KELB":
		c.Redirect(http.StatusTemporaryRedirect, "/kelb")
	case "PPK":
		c.Redirect(http.StatusTemporaryRedirect, "/ppk")
	case "PPE":
		c.Redirect(http.StatusTemporaryRedirect, "/ppe")
	case "ULP":
		c.Redirect(http.StatusTemporaryRedirect, "/ulp")
	case "Super-Admin":
		c.Redirect(http.StatusTemporaryRedirect, "/super-admin")
	}
}

func SPKProgress(c *gin.Context) {
	c.HTML(http.StatusOK, "spk-progress.html", struct {
		ID string
	}{
		ID: c.Params.ByName("id"),
	})
}

func ChangePWD(c *gin.Context) {
	_, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		c.JSON(http.StatusForbidden, "forbidden")
		return
	}
	c.HTML(http.StatusOK, "change-pwd.html", nil)
}

func Revision(c *gin.Context) {
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		c.JSON(http.StatusForbidden, "forbidden")
		return
	}
	if entity.Role != "PPK" {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	c.HTML(http.StatusOK, "revision.html", struct {
		ID    string
		SPKID string
	}{
		ID:    c.Params.ByName("id"),
		SPKID: c.Params.ByName("spk_id"),
	})
}

func RejectPPP(c *gin.Context) {
	c.HTML(http.StatusOK, "reject-ppp.html", struct {
		PPPID   string
		INBOXID string
	}{
		PPPID:   c.Params.ByName("ppp_id"),
		INBOXID: c.Params.ByName("inbox_id"),
	})
}

func RejectRP(c *gin.Context) {
	c.HTML(http.StatusOK, "reject-rp.html", struct {
		RPID    string
		INBOXID string
	}{
		RPID:    c.Params.ByName("rp_id"),
		INBOXID: c.Params.ByName("inbox_id"),
	})
}
