package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"manajemen-work-order/models"
	"manajemen-work-order/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const NOIMGPATH = "archive/img/noimg.png"

func PPPGet(c *gin.Context) {
	//validate entity that entity role is super-admin
	_, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, err.Error())
		return
	}
	mdl := models.PPP{}
	//get last-id
	val, _ := c.GetQuery("last-id")
	var lastID int64

	if val == "" {
		lastID = 0
	} else {
		valInt, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			services.SendBasicResponse(c, http.StatusBadRequest, false, err.Error())
			return
		}
		lastID = valInt
	}

	//extract db
	ctx := context.Background()
	db, err := services.GetDB(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}
	res, err := mdl.FindAll(db, ctx, lastID)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

func PPPPost(c *gin.Context) {

	//validate entity
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, err.Error())
		return
	}

	//decode ppp
	var payload models.PPP

	nota := c.PostForm("nota")
	pekerjaan := c.PostForm("pekerjaan")
	perihal := c.PostForm("perihal")
	sifat := c.PostForm("sifat")
	photo, err := c.FormFile("photo")

	//if there is photo payload, then store and set payload.Photo = pathImage if not set payload.Photo to NOIMG
	if err != nil {
		payload.Photo = NOIMGPATH
	} else {
		pathImage := fmt.Sprintf("archive/img/%s%s%s", entity.Fullname, time.Now(), photo.Filename)

		err = c.SaveUploadedFile(photo, pathImage)
		if err != nil {
			services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
			return
		}

		payload.Photo = pathImage
	}

	payload.Nota = nota
	payload.Pekerjaan = pekerjaan
	payload.Perihal = perihal
	payload.Sifat = sifat

	//TOBE generate pdf and store it on disk
	pathDoc, err := services.CreatePDFofPPP(payload, entity, "BDMU")
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}
	payload.Doc = pathDoc
	payload.Status = "ON ROUTE TO BDMU"

	//validation
	err = services.IsNotEmpty(payload.Doc, payload.Sifat, payload.Status, payload.Nota, payload.Perihal, payload.Pekerjaan, payload.Photo)
	if err != nil {
		if payload.Photo != NOIMGPATH {
			services.RemoveFile(payload.Photo)
		}
		services.RemoveFile(pathDoc)
		services.SendBasicResponse(c, http.StatusBadRequest, false, err.Error())
		return
	}

	//extract db
	db, err := services.GetDB(c)
	if err != nil {
		if payload.Photo != NOIMGPATH {
			services.RemoveFile(payload.Photo)
		}
		services.RemoveFile(pathDoc)
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//store information to db (ppp and bdmu_ppp)
	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		if payload.Photo != NOIMGPATH {
			services.RemoveFile(payload.Photo)
		}
		services.RemoveFile(pathDoc)
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	res, err := payload.InsertTx(tx, ctx, entity.ID)
	if err != nil {
		if payload.Photo != NOIMGPATH {
			services.RemoveFile(payload.Photo)
		}
		services.RemoveFile(pathDoc)
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	insertedID, err := res.LastInsertId()
	if err != nil {
		if payload.Photo != NOIMGPATH {
			services.RemoveFile(payload.Photo)
		}
		services.RemoveFile(pathDoc)
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
		if payload.Photo != NOIMGPATH {
			services.RemoveFile(payload.Photo)
		}
		services.RemoveFile(pathDoc)
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	err = tx.Commit()
	if err != nil {
		if payload.Photo != NOIMGPATH {
			services.RemoveFile(payload.Photo)
		}
		services.RemoveFile(pathDoc)
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	services.SendBasicResponse(c, http.StatusOK, true, "PPP created and routed to BDMU")
}

func PPPOKBDMU(c *gin.Context) {
	//validate entity must be bdmu
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, err.Error())
		return
	}

	if entity.Role != "BDMU" {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, "NOT BDMU")
		return
	}

	//extract id from param
	val := c.Params.ByName("ppp_id")
	pppid, err := strconv.ParseInt(val, 10, 64)
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

	//change status and input bdmu_id on ppp
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		tx.Rollback()
		return
	}

	ppp := models.PPP{
		ID:     pppid,
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
		PPPID: pppid,
	}
	_, err = bdmupppp.InsertTx(tx, ctx)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		tx.Rollback()
		return
	}

	//delete bdmu_ppp
	bdmuppp := models.BDMUPPP{
		ID: bdmuid,
	}
	_, err = bdmuppp.DeleteTx(tx, ctx)
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

func PPPOKBDMUP(c *gin.Context) {
	//validate entity must be bdmu
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, err.Error())
		return
	}

	if entity.Role != "BDMUP" {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, "NOT BDMUP")
		return
	}

	//extract id from param
	val := c.Params.ByName("ppp_id")
	pppid, err := strconv.ParseInt(val, 10, 64)
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

	ppp := models.PPP{
		ID:      pppid,
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
		PPPID: pppid,
	}
	_, err = kelappp.InsertTx(tx, ctx)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		tx.Rollback()
		return
	}

	//delete bdmup_ppp
	bdmupppp := models.BDMUPPPP{
		ID: bdmupid,
	}

	_, err = bdmupppp.DeleteTx(tx, ctx)
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

func PPPOKKELA(c *gin.Context) {
	//validate entity must be bdmu
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, err.Error())
		return
	}

	if entity.Role != "KELA" {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, "NOT KELA")
		return
	}

	//extract id from param
	val := c.Params.ByName("ppp_id")
	pppid, err := strconv.ParseInt(val, 10, 64)
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

	ppp := models.PPP{
		ID:     pppid,
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
		PPPID: pppid,
	}
	_, err = kelbppp.InsertTx(tx, ctx)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		tx.Rollback()
		return
	}

	//delete kela_ppp
	kelappp := models.KELAPPP{
		ID: kelaid,
	}
	_, err = kelappp.DeleteTx(tx, ctx)
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

func PPPNO(c *gin.Context) {
	//validate entity must be bdmu || bdmup || kela
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, err.Error())
		return
	}

	if entity.Role != "BDMU" && entity.Role != "BDMUP" && entity.Role != "KELA" {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, "NOT BDMU")
		return
	}

	//extract ppp_id and inboxid from param
	val := c.Params.ByName("ppp_id")
	pppid, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}
	val2 := c.Params.ByName("inbox_id")
	inboxid, err := strconv.ParseInt(val2, 10, 64)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//decode payload
	ppp := models.PPP{}

	err = json.NewDecoder(c.Request.Body).Decode(&ppp)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//validation
	err = services.IsNotEmpty(ppp.Reason)
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

	//change status ppp

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	ppp.ID = pppid
	ppp.Status = "REJECTED"

	_, err = ppp.UpdateStatusAndReasonTx(tx, ctx)
	if err != nil {
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//delete inbox
	switch entity.Role {
	case "BDMU":
		mdl := models.BDMUPPP{
			ID: inboxid,
		}
		_, err = mdl.DeleteTx(tx, ctx)
	case "BDMUP":
		mdl := models.BDMUPPPP{
			ID: inboxid,
		}
		_, err = mdl.DeleteTx(tx, ctx)
	case "KELA":
		mdl := models.KELAPPP{
			ID: inboxid,
		}
		_, err = mdl.DeleteTx(tx, ctx)
	}
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

	//send respond
	services.SendBasicResponse(c, http.StatusOK, true, "ppp rejected")
}
