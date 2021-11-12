package controllers

import (
	"context"
	"manajemen-work-order/models"
	"manajemen-work-order/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func KELBPPPDelete(c *gin.Context) {
	//validate entity that entity role is super-admin
	entity, err := services.ValidateTokenFromHeader(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, err.Error())
		return
	}

	if entity.Role != "KELB" {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, "NOT Super-Admin")
		return
	}
	//extract ppp_id from param
	val := c.Params.ByName("id")
	id, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//extract db
	ctx := context.Background()
	db, err := services.GetDB(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	mdl := models.KELBPPP{
		ID: id,
	}

	_, err = mdl.Delete(db, ctx)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	services.SendBasicResponse(c, http.StatusOK, true, "delete success")
}

func KELBPPPGet(c *gin.Context) {

}
