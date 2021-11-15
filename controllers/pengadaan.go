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

func PengadaanGet(c *gin.Context) {
	//validate entity that entity role is super-admin
	_, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, err.Error())
		return
	}
	mdl := models.Pengadaan{}
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

func PengadaanFromPPE(c *gin.Context) {
	pengadaanHelper(c, "PPE")
}

func PengadaanFromULP(c *gin.Context) {
	pengadaanHelper(c, "ULP")
}

func pengadaanHelper(c *gin.Context, role string) {
	//validate entity must be ulp / ppe
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, err.Error())
		return
	}

	if entity.Role != role {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, "NOT ULP / PPE")
		return
	}

	//get perkiraan_biaya_id dan inbox_id from params
	val := c.Params.ByName("perkiraan_biaya_id")
	perkiraanbiayaid, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		services.SendBasicResponse(c, http.StatusBadRequest, false, err.Error())
		return
	}
	//get inbox_id from params
	val2 := c.Params.ByName("inbox_id")
	inboxid, err := strconv.ParseInt(val2, 10, 64)
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

	docPath := fmt.Sprintf("archive/pengadaan/%s%s%s", time.Now(), entity.Fullname, doc.Filename)

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

	//store to pengaadan
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		services.RemoveFile(docPath)
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	pengaadaan := models.Pengadaan{
		PerkiraanBiayaID: perkiraanbiayaid,
		Doc:              docPath,
	}

	res, err := pengaadaan.InsertTx(tx, ctx, entity.ID)
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

	//store to ppk_pengadaan
	ppkpengadaan := models.PPKPengadaan{
		PengadaanID: lastInsertedID,
	}
	_, err = ppkpengadaan.InsertTx(tx, ctx)
	if err != nil {
		services.RemoveFile(docPath)
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//delete ulp or ppe perkiraan_biaya
	switch role {
	case "PPE":
		ppeperkiaanbiaya := models.PPEPerkiraanBiaya{
			ID: inboxid,
		}
		_, err = ppeperkiaanbiaya.DeleteTx(tx, ctx)
	case "ULP":
		ulpperkiraanbiaya := models.ULPPerkiraanBiaya{
			ID: inboxid,
		}
		_, err = ulpperkiraanbiaya.DeleteTx(tx, ctx)
	}
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

	//respond
	services.SendBasicResponse(c, http.StatusOK, true, "pengadaan created")
}
