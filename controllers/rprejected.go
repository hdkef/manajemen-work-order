package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"manajemen-work-order/models"
	"manajemen-work-order/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RPRejectedPost(c *gin.Context) {
	//validate entity must be kela, bdmup, or bdmu
	entity, err := services.ValidateTokenFromHeader(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, err.Error())
		return
	}

	if entity.Role != "BDMU" && entity.Role != "BDMUP" && entity.Role != "KELA" {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, "unauthorized")

		return
	}

	//get rp_id from params
	val := c.Params.ByName("rp_id")
	rpid, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		services.SendBasicResponse(c, http.StatusBadRequest, false, err.Error())
		return
	}

	//decode payload
	RPRejected := models.RPRejected{}

	err = json.NewDecoder(c.Request.Body).Decode(&RPRejected)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//validate
	err = services.IsNotEmpty(RPRejected.MSG)
	if err != nil {
		services.SendBasicResponse(c, http.StatusBadRequest, false, err.Error())
		return
	}

	//extract db
	db, err := services.GetDB(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//store to rp_rejected
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	RPRejected.RPID = rpid

	_, err = RPRejected.InsertTx(tx, ctx, entity.ID)
	if err != nil {
		fmt.Println("entityid", entity.ID)
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//change status of RP
	rp := models.RP{
		ID:     rpid,
		Status: "REJECTED PLS REMAKE",
	}
	_, err = rp.UpdateStatusTx(tx, ctx)
	if err != nil {
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//send
	services.SendBasicResponse(c, http.StatusOK, true, "RP rejected.")
}
