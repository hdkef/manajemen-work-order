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
	pin := rand.Int63()

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
	fmt.Println(emailMsg)
	err = services.SendEmail(workerEmail, "Email dari Server", emailMsg)
	if err != nil {
		services.SendBasicResponse(c, http.StatusOK, true, fmt.Sprintf("spk is created but failed to send. Please send email manually (%s) work order id %d with pin %d", workerEmail, lastInsertedID, pin))
		return
	}
	//respond
	services.SendBasicResponse(c, http.StatusOK, true, "spk is created and sent to worker")

}

func SPKLapor(c *gin.Context) {

}

func SPKOK(c *gin.Context) {

}

func SPKNO(c *gin.Context) {

}
