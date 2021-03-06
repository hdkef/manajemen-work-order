package controllers

import (
	"context"
	"manajemen-work-order/models"
	"manajemen-work-order/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func BDMUPRPDelete(c *gin.Context) {
// 	//validate entity that entity role is super-admin
// 	entity, err := services.ValidateTokenFromCookie(c)
// 	if err != nil {
// 		services.SendBasicResponse(c, http.StatusUnauthorized, false, err.Error())
// 		return
// 	}

// 	if entity.Role != "BDMUP" {
// 		services.SendBasicResponse(c, http.StatusUnauthorized, false, "NOT Super-Admin")
// 		return
// 	}
// 	//extract ppp_id from param
// 	val := c.Params.ByName("id")
// 	id, err := strconv.ParseInt(val, 10, 64)
// 	if err != nil {
// 		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
// 		return
// 	}

// 	//extract db
// 	ctx := context.Background()
// 	db, err := services.GetDB(c)
// 	if err != nil {
// 		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
// 		return
// 	}

// 	mdl := models.BDMUPRP{
// 		ID: id,
// 	}

// 	_, err = mdl.Delete(db, ctx)
// 	if err != nil {
// 		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
// 		return
// 	}

// 	services.SendBasicResponse(c, http.StatusOK, true, "delete success")
// }

func BDMUPRPGet(c *gin.Context) {
	//validate entity that entity role is super-admin
	_, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, err.Error())
		return
	}
	mdl := models.BDMUPRP{}
	//extract db
	ctx := context.Background()
	db, err := services.GetDB(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}
	res, err := mdl.FindAll(db, ctx)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}
