package controllers

import (
	"context"
	"fmt"
	"manajemen-work-order/models"
	"manajemen-work-order/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func RP(c *gin.Context) {
	//validate entity must be kelb
	entity, err := services.ValidateTokenFromHeader(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, err.Error())
		return
	}

	if entity.Role != "KELB" {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, "NOT KELB")
		return
	}

	//get ppp_id from params
	val := c.Params.ByName("ppp_id")
	pppid, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		services.SendBasicResponse(c, http.StatusBadRequest, false, err.Error())
		return
	}
	//get kelb_ppp_id from params
	val2 := c.Params.ByName("kelb_ppp_id")
	kelbpppid, err := strconv.ParseInt(val2, 10, 64)
	if err != nil {
		services.SendBasicResponse(c, http.StatusBadRequest, false, err.Error())
		return
	}

	//decode payload

	doc, err := c.FormFile("doc")
	if err != nil {
		services.SendBasicResponse(c, http.StatusBadRequest, false, err.Error())
		return
	}

	docPath := fmt.Sprintf("archive/rp/%s%s", entity.Fullname, time.Now())

	err = c.SaveUploadedFile(doc, docPath)
	if err != nil {
		services.RemoveFile(docPath)
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//extract db
	db, err := services.GetDB(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//store to rp
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	rp := models.RP{
		Doc:    docPath,
		Status: "ON ROUTE TO KELA",
		PPPID:  pppid,
	}
	res, err := rp.InsertTx(tx, ctx, entity.ID)
	if err != nil {
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//store to kela_rp
	kelarp := models.KELARP{
		RPID: lastInsertedID,
	}
	_, err = kelarp.InsertTx(tx, ctx)
	if err != nil {
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//delete kelb_ppp
	kelbppp := models.KELBPPP{
		ID: kelbpppid,
	}
	_, err = kelbppp.DeleteTx(tx, ctx)
	if err != nil {
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//change status of ppp
	ppp := models.PPP{
		ID: pppid,
	}
	_, err = ppp.UpdateStatusTx(tx, ctx)
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

	services.SendBasicResponse(c, http.StatusOK, true, "RP created and routed to KELA")

}

func RPOKKELA(c *gin.Context) {

}

func RPNOKELA(c *gin.Context) {

}

func RPOKBDMU(c *gin.Context) {

}

func RPNOBDMU(c *gin.Context) {

}

func RPOKBDMUP(c *gin.Context) {

}

func RPNOBDMUP(c *gin.Context) {

}
