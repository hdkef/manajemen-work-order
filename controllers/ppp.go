package controllers

import (
	"context"
	"encoding/json"
	"manajemen-work-order/models"
	"manajemen-work-order/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PPPPost(c *gin.Context) {

	//validate entity
	entity, err := services.ValidateTokenFromHeader(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, err.Error())
		return
	}

	//decode ppp
	var payload models.PPP
	err = json.NewDecoder(c.Request.Body).Decode(&payload)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//generate pdf and store it on disk
	path, err := services.CreatePDFofPPP(payload, entity, "BDMU")
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}
	payload.Doc = path
	payload.Status = "ON ROUTE TO BDMU"

	//validation
	err = services.IsNotEmpty(payload.Doc, payload.Sifat, payload.Status, payload.Nota, payload.Perihal, payload.Pekerjaan)
	if err != nil {
		services.RemoveFile(path)
		services.SendBasicResponse(c, http.StatusBadRequest, false, err.Error())
		return
	}

	//extract db
	db, err := services.GetDB(c)
	if err != nil {
		services.RemoveFile(path)
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//store information to db (ppp and bdmu_ppp)
	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		services.RemoveFile(path)
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	res, err := payload.InsertTx(tx, ctx, entity.ID, 0, 0, 0)
	if err != nil {
		services.RemoveFile(path)
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	insertedID, err := res.LastInsertId()
	if err != nil {
		services.RemoveFile(path)
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//store information to bdmu_ppp
	bdmuppp := models.BDMUPPP{
		PPPID: insertedID,
	}

	_, err = bdmuppp.InsertTx(tx, ctx)
	if err != nil {
		services.RemoveFile(path)
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	err = tx.Commit()
	if err != nil {
		services.RemoveFile(path)
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	services.SendBasicResponse(c, http.StatusOK, true, "PPP created and routed to BDMU")
}

func PPPOKBDMU(c *gin.Context) {
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
	val := c.Params.ByName("id")
	id, err := strconv.ParseInt(val, 10, 64)
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

	//change status and input bdmu_id on ppp
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		tx.Rollback()
		return
	}

	ppp := models.PPP{
		ID:     id,
		Status: "ON ROUTE TO BDMUP",
		BDMUID: entity.ID,
	}

	_, err = ppp.UpdateStatusAndBDMUIDTx(tx, ctx)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		tx.Rollback()
		return
	}

	//store to dbmup_ppp
	bdmupppp := models.BDMUPPPP{
		PPPID: id,
	}
	_, err = bdmupppp.InsertTx(tx, ctx)
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
	services.SendBasicResponse(c, http.StatusOK, true, "PPP routed to BDMUP")
}

func PPPNOBDMU(c *gin.Context) {

}

func PPPOKBDMUP(c *gin.Context) {
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
	val := c.Params.ByName("id")
	id, err := strconv.ParseInt(val, 10, 64)
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

	ppp := models.PPP{
		ID:      id,
		Status:  "ON ROUTE TO KELA",
		BDMUPID: entity.ID,
	}

	_, err = ppp.UpdateStatusAndBDMUPIDTx(tx, ctx)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		tx.Rollback()
		return
	}

	//store to kela_ppp
	kelappp := models.KELAPPP{
		PPPID: id,
	}
	_, err = kelappp.InsertTx(tx, ctx)
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
	services.SendBasicResponse(c, http.StatusOK, true, "PPP routed to KELA")
}

func PPPNOBDMUP(c *gin.Context) {

}

func PPPOKKELA(c *gin.Context) {
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
	val := c.Params.ByName("id")
	id, err := strconv.ParseInt(val, 10, 64)
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

	ppp := models.PPP{
		ID:     id,
		Status: "ON ROUTE TO KELB",
		KELAID: entity.ID,
	}

	_, err = ppp.UpdateStatusAndKELAIDTx(tx, ctx)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		tx.Rollback()
		return
	}

	//store to kelb_ppp
	kelbppp := models.KELBPPP{
		PPPID: id,
	}
	_, err = kelbppp.InsertTx(tx, ctx)
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
	services.SendBasicResponse(c, http.StatusOK, true, "PPP routed to KELB")
}

func PPPNOKELA(c *gin.Context) {

}
