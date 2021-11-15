package views

import (
	"context"
	"manajemen-work-order/models"
	"manajemen-work-order/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PPPAll(c *gin.Context) {
	//auth
	_, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	//extract from db
	ctx := context.Background()
	db, err := services.GetDB(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//get data
	mdl := models.PPP{}
	data, err := mdl.FindAll(db, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//send data and render
	c.HTML(http.StatusOK, "ppp-all.html", data)
}

func RPAll(c *gin.Context) {
	//auth
	_, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	//extract from db
	ctx := context.Background()
	db, err := services.GetDB(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//get data
	mdl := models.RP{}
	data, err := mdl.FindAll(db, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//send data and render
	c.HTML(http.StatusOK, "rp-all.html", data)
}

func PerkiraanBiayaAll(c *gin.Context) {
	//auth
	_, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	//extract from db
	ctx := context.Background()
	db, err := services.GetDB(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//get data
	mdl := models.PerkiraanBiaya{}
	data, err := mdl.FindAll(db, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//send data and render
	c.HTML(http.StatusOK, "perkiraan-biaya-all.html", data)
}

func PengadaanAll(c *gin.Context) {
	//auth
	_, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	//extract from db
	ctx := context.Background()
	db, err := services.GetDB(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//get data
	mdl := models.Pengadaan{}
	data, err := mdl.FindAll(db, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//send data and render
	c.HTML(http.StatusOK, "pengadaan-all.html", data)
}

func SPKAll(c *gin.Context) {
	//auth
	_, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	//extract from db
	ctx := context.Background()
	db, err := services.GetDB(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//get data
	mdl := models.SPK{}
	data, err := mdl.FindAll(db, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//send data and render
	c.HTML(http.StatusOK, "spk-all.html", data)
}
