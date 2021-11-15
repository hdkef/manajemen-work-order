package views

import (
	"context"
	"manajemen-work-order/models"
	"manajemen-work-order/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func EntityOne(c *gin.Context) {
	//auth
	_, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	//get id from path
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
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
	mdl := models.Entity{
		ID: id,
	}
	err = mdl.FindOneByID(db, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//send data and render
	c.HTML(http.StatusOK, "entity-one.html", mdl)
}

func PPPOne(c *gin.Context) {
	//auth
	_, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	//get id from path
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
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
	mdl := models.PPP{
		ID: id,
	}
	data, err := mdl.FindOne(db, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//send data and render
	c.HTML(http.StatusOK, "ppp-one.html", data)
}

func RPOne(c *gin.Context) {
	//auth
	_, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	//get id from path
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
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
	mdl := models.RP{
		ID: id,
	}
	data, err := mdl.FindOne(db, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//send data and render
	c.HTML(http.StatusOK, "rp-one.html", data)
}

func PerkiraanBiayaOne(c *gin.Context) {
	//auth
	_, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	//get id from path
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
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
	mdl := models.PerkiraanBiaya{
		ID: id,
	}
	data, err := mdl.FindOne(db, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//send data and render
	c.HTML(http.StatusOK, "perkiraan-biaya-one.html", data)
}

func PengadaanOne(c *gin.Context) {
	//auth
	_, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	//get id from path
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
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
	mdl := models.Pengadaan{
		ID: id,
	}
	data, err := mdl.FindOne(db, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//send data and render
	c.HTML(http.StatusOK, "pengadaan-one.html", data)
}

func SPKOne(c *gin.Context) {
	//auth
	_, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	//get id from path
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
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
	mdl := models.SPK{
		ID: id,
	}
	data, err := mdl.FindOne(db, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//send data and render
	c.HTML(http.StatusOK, "spk-one.html", data)
}
