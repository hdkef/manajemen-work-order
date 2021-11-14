package controllers

import (
	"context"
	"fmt"
	"manajemen-work-order/models"
	"manajemen-work-order/services"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SPKGet(c *gin.Context) {
	//validate entity that entity role is super-admin
	_, err := services.ValidateTokenFromHeader(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, err.Error())
		return
	}
	mdl := models.SPK{}
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

func SPKPost(c *gin.Context) {
	//validate entity must be ppk
	entity, err := services.ValidateTokenFromHeader(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, err.Error())
		return
	}

	if entity.Role != "PPK" {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, "NOT PPK")
		return
	}

	//get pengadaan_id dan ppk_pengadaan_id from params
	val := c.Params.ByName("pengadaan_id")
	pengadaanid, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		services.SendBasicResponse(c, http.StatusBadRequest, false, err.Error())
		return
	}
	//get ppk_pengadaan_id from params
	val2 := c.Params.ByName("ppk_pengadaan_id")
	ppkpengadaanid, err := strconv.ParseInt(val2, 10, 64)
	if err != nil {
		services.SendBasicResponse(c, http.StatusBadRequest, false, err.Error())
		return
	}

	//decode payload

	workerEmail := c.PostForm("worker_email")
	//validate email
	err = services.IsNotEmpty(workerEmail)
	if err != nil {
		services.SendBasicResponse(c, http.StatusBadRequest, false, err.Error())
		return
	}
	err = services.IsEmail(workerEmail)
	if err != nil {
		services.SendBasicResponse(c, http.StatusBadRequest, false, err.Error())
		return
	}

	doc, err := c.FormFile("doc")
	if err != nil {
		services.SendBasicResponse(c, http.StatusBadRequest, false, err.Error())
		return
	}

	docPath := fmt.Sprintf("archive/spk/%s%s", entity.Fullname, time.Now())

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

	//store to spk
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		services.RemoveFile(docPath)
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	spk := models.SPK{
		Doc:         docPath,
		PengadaanID: pengadaanid,
		Status:      "Created",
		WorkerEmail: workerEmail,
	}
	res, err := spk.InsertTx(tx, ctx, entity.ID)
	if err != nil {
		services.RemoveFile(docPath)
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//delete ppk pengadaan
	ppkpengadaan := models.PPKPengadaan{
		ID: ppkpengadaanid,
	}
	_, err = ppkpengadaan.DeleteTx(tx, ctx)
	if err != nil {
		services.RemoveFile(docPath)
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//create email session
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		services.RemoveFile(docPath)
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//create random int
	pin := int64(rand.Int31())

	emailSession := models.EmailSession{
		SPKID: lastInsertedID,
		PIN:   pin,
	}
	_, err = emailSession.InsertTx(tx, ctx)
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

	//TOBE send email to worker
	emailMsg := fmt.Sprintf("work order has been created click this link to see details. Please submit work progression to http://localhost:8080/spk/progress and use work order id of %d and pin %d", lastInsertedID, pin)

	err = services.SendEmail(workerEmail, "Email dari Server", emailMsg)
	if err != nil {
		services.SendBasicResponse(c, http.StatusOK, true, fmt.Sprintf("spk is created but failed to send. Please send email manually (%s) work order id %d with pin %d", workerEmail, lastInsertedID, pin))
		return
	}
	//respond
	services.SendBasicResponse(c, http.StatusOK, true, "spk is created and sent to worker")

}

func SPKEdit(c *gin.Context) {
	//get spk_id from params
	idString := c.Params.ByName("id")
	spkid, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		services.SendBasicResponse(c, http.StatusBadRequest, true, err.Error())
		return
	}

	//decode payload
	status := c.PostForm("status")
	pinString := c.PostForm("pin")
	pin, err := strconv.ParseInt(pinString, 10, 64)
	if err != nil {
		services.SendBasicResponse(c, http.StatusBadRequest, true, err.Error())
		return
	}

	//validate
	err = services.IsNotEmpty(status)
	if err != nil {
		services.SendBasicResponse(c, http.StatusBadRequest, true, err.Error())
		return
	}

	//extract db
	db, err := services.GetDB(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//get pin from db
	ctx := context.Background()
	emailSessionMdl := models.EmailSession{
		SPKID: spkid,
	}

	emailSessionFromDB, err := emailSessionMdl.FindOneBySPKID(db, ctx)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//compare pin
	if emailSessionFromDB.PIN != pin {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, "wrong pin")
		return
	}

	//update status spk
	spk := models.SPK{
		ID:     spkid,
		Status: status,
	}
	_, err = spk.UpdateStatus(db, ctx)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	services.SendBasicResponse(c, http.StatusOK, true, "spk progress updated")

}

func SPKLapor(c *gin.Context) {
	//get spk_id from params
	idString := c.Params.ByName("id")
	spkid, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		services.SendBasicResponse(c, http.StatusBadRequest, true, err.Error())
		return
	}

	//decode payload
	pinString := c.PostForm("pin")
	pin, err := strconv.ParseInt(pinString, 10, 64)
	if err != nil {
		services.SendBasicResponse(c, http.StatusBadRequest, true, err.Error())
		return
	}

	//extract db
	db, err := services.GetDB(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	emailSessionMdl := models.EmailSession{
		SPKID: spkid,
	}

	emailSessionFromDB, err := emailSessionMdl.FindOneBySPKID(db, ctx)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//compare pin
	if emailSessionFromDB.PIN != pin {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, "wrong pin")
		return
	}

	//update status of spk
	spk := models.SPK{
		ID:     spkid,
		Status: "spk closed by worker",
	}

	_, err = spk.UpdateStatusTx(tx, ctx)
	if err != nil {
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//store to ppk_spk
	ppkspk := models.PPKSPK{
		SPKID: spkid,
	}

	_, err = ppkspk.InsertTx(tx, ctx)
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
	services.SendBasicResponse(c, http.StatusOK, true, "spk is closed by worker. Will be checked soon.")
}

func SPKOK(c *gin.Context) {
	//validate entity must be ppk
	entity, err := services.ValidateTokenFromHeader(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, err.Error())
		return
	}

	if entity.Role != "PPK" {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, "NOT PPK")
		return
	}

	//get spkid from params
	val := c.Params.ByName("id")
	spkid, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		services.SendBasicResponse(c, http.StatusBadRequest, false, err.Error())
		return
	}
	//get kelb_ppp_id from params
	val2 := c.Params.ByName("ppk_spk_id")
	ppkspkid, err := strconv.ParseInt(val2, 10, 64)
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

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//change status spk
	spk := models.SPK{
		ID:     spkid,
		Status: "closed by PPK",
	}

	_, err = spk.UpdateStatusTx(tx, ctx)
	if err != nil {
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//get worker_email from spk
	workerEmail, err := spk.FindWorkerEmailTx(tx, ctx)
	if err != nil {
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//delete ppkspk
	ppkspk := models.PPKSPK{
		ID: ppkspkid,
	}

	_, err = ppkspk.DeleteTx(tx, ctx)
	if err != nil {
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//delete email session
	emailSession := models.EmailSession{
		SPKID: spkid,
	}
	_, err = emailSession.DeleteBySPKIDTx(tx, ctx)
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

	//send email
	err = services.SendEmail(workerEmail, "Email dari Server", fmt.Sprintf("work order dengan id %d, telah ditutup oleh PPK", spkid))
	if err != nil {
		services.SendBasicResponse(c, http.StatusOK, true, fmt.Sprintf("spk is closed by PPK but email respond has not been sent to worker (%s). Please send it manually", workerEmail))
		return
	}

	//send
	services.SendBasicResponse(c, http.StatusOK, true, "spk is closed by PPK")
}

func SPKNO(c *gin.Context) {
	//validate entity must be ppk
	entity, err := services.ValidateTokenFromHeader(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, err.Error())
		return
	}

	if entity.Role != "PPK" {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, "NOT PPK")
		return
	}

	//get spkid from params
	val := c.Params.ByName("id")
	spkid, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		services.SendBasicResponse(c, http.StatusBadRequest, false, err.Error())
		return
	}
	//get kelb_ppp_id from params
	val2 := c.Params.ByName("ppk_spk_id")
	ppkspkid, err := strconv.ParseInt(val2, 10, 64)
	if err != nil {
		services.SendBasicResponse(c, http.StatusBadRequest, false, err.Error())
		return
	}

	//decode payload
	msg := c.PostForm("msg")

	err = services.IsNotEmpty(msg)
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

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//change status spk
	spk := models.SPK{
		ID:     spkid,
		Status: "Need Revision.",
	}

	_, err = spk.UpdateStatusTx(tx, ctx)
	if err != nil {
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//get worker_email from spk
	workerEmail, err := spk.FindWorkerEmailTx(tx, ctx)
	if err != nil {
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//delete ppkspk
	ppkspk := models.PPKSPK{
		ID: ppkspkid,
	}

	_, err = ppkspk.DeleteTx(tx, ctx)
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

	//send email
	err = services.SendEmail(workerEmail, "Email dari Server", fmt.Sprintf("work order dengan id %d rejected dengan pesan ==> %s", spkid, msg))
	if err != nil {
		services.SendBasicResponse(c, http.StatusOK, true, fmt.Sprintf("SPK is rejected and need revision but email respond has not been sent to worker (%s). Please send it manually", workerEmail))
		return
	}

	//send
	services.SendBasicResponse(c, http.StatusOK, true, "spk is rejected and need revision")
}
