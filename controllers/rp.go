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
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//extract db
	db, err := services.GetDB(c)
	if err != nil {
		services.RemoveFile(docPath)
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//store to rp
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		services.RemoveFile(docPath)
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
		services.RemoveFile(docPath)
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		services.RemoveFile(docPath)
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
		services.RemoveFile(docPath)
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
		services.RemoveFile(docPath)
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//change status of ppp
	ppp := models.PPP{
		ID:     pppid,
		Status: "RP is created",
	}
	_, err = ppp.UpdateStatusTx(tx, ctx)
	if err != nil {
		services.RemoveFile(docPath)
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	err = tx.Commit()
	if err != nil {
		services.RemoveFile(docPath)
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	services.SendBasicResponse(c, http.StatusOK, true, "RP created and routed to KELA")

}

func RPOKKELA(c *gin.Context) {
	//validate entity must be bdmu
	entity, err := services.ValidateTokenFromHeader(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, err.Error())
		return
	}

	if entity.Role != "KELA" {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, "NOT KELA")
		return
	}

	//extract id from param
	val := c.Params.ByName("rp_id")
	rpid, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}
	val2 := c.Params.ByName("kela_id")
	kelaid, err := strconv.ParseInt(val2, 10, 64)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//extract db
	db, err := services.GetDB(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//change status and input bdmup_id on ppp
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		tx.Rollback()
		return
	}

	rp := models.RP{
		ID:     rpid,
		Status: "ON ROUTE TO BDMUP",
		KELAID: entity.ID,
	}

	_, err = rp.UpdateStatusAndKELAIDTx(tx, ctx)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		tx.Rollback()
		return
	}

	//store to bdmup_rp
	bdmuprp := models.BDMUPRP{
		RPID: rpid,
	}
	_, err = bdmuprp.InsertTx(tx, ctx)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		tx.Rollback()
		return
	}

	//delete kela_rp
	kelarp := models.KELARP{
		ID: kelaid,
	}

	_, err = kelarp.DeleteTx(tx, ctx)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		tx.Rollback()
		return
	}

	//response
	services.SendBasicResponse(c, http.StatusOK, true, "RP routed to BDMUP")
}

func RPNOKELA(c *gin.Context) {

}

func RPOKBDMUP(c *gin.Context) {
	//validate entity must be bdmu
	entity, err := services.ValidateTokenFromHeader(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, err.Error())
		return
	}

	if entity.Role != "BDMUP" {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, "NOT BDMUP")
		return
	}

	//extract id from param
	val := c.Params.ByName("rp_id")
	rpid, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}
	val2 := c.Params.ByName("bdmup_id")
	bdmupid, err := strconv.ParseInt(val2, 10, 64)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//extract db
	db, err := services.GetDB(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//change status and input bdmup_id on ppp
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		tx.Rollback()
		return
	}

	rp := models.RP{
		ID:     rpid,
		Status: "ON ROUTE TO BDMU",
		KELAID: entity.ID,
	}

	_, err = rp.UpdateStatusAndBDMUPIDTx(tx, ctx)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		tx.Rollback()
		return
	}

	//store to bdmu_rp
	bdmuprp := models.BDMUPRP{
		RPID: rpid,
	}
	_, err = bdmuprp.InsertTx(tx, ctx)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		tx.Rollback()
		return
	}

	//delete bdmup rp
	bdmup := models.BDMUPRP{
		ID: bdmupid,
	}

	_, err = bdmup.DeleteTx(tx, ctx)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		tx.Rollback()
		return
	}

	//response
	services.SendBasicResponse(c, http.StatusOK, true, "RP routed to BDMU")
}

func RPNOBDMUP(c *gin.Context) {

}

func RPOKBDMU(c *gin.Context) {
	//validate entity must be bdmu
	entity, err := services.ValidateTokenFromHeader(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, err.Error())
		return
	}

	if entity.Role != "BDMU" {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, "NOT BDMU")
		return
	}

	//extract id from param
	val := c.Params.ByName("rp_id")
	rpid, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}
	val2 := c.Params.ByName("bdmu_id")
	bdmuid, err := strconv.ParseInt(val2, 10, 64)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//extract db
	db, err := services.GetDB(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//change status and input bdmup_id on ppp
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		tx.Rollback()
		return
	}

	rp := models.RP{
		ID:     rpid,
		Status: "ON ROUTE TO PPK",
		KELAID: entity.ID,
	}

	_, err = rp.UpdateStatusAndBDMUIDTx(tx, ctx)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		tx.Rollback()
		return
	}

	//store to ppk_rp
	ppkrp := models.PPKRP{
		RPID: rpid,
	}
	_, err = ppkrp.InsertTx(tx, ctx)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		tx.Rollback()
		return
	}

	//delete bdmu rp
	bdmurp := models.BDMUPRP{
		ID: bdmuid,
	}

	_, err = bdmurp.DeleteTx(tx, ctx)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		tx.Rollback()
		return
	}

	//response
	services.SendBasicResponse(c, http.StatusOK, true, "RP routed to PPK")
}

func RPNOBDMU(c *gin.Context) {

}
