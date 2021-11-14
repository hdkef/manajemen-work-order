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

func PerkiraanBiayaGet(c *gin.Context) {
	//validate entity that entity role is super-admin
	_, err := services.ValidateTokenFromHeader(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, err.Error())
		return
	}
	mdl := models.PerkiraanBiaya{}
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

func ULPPerkiraanBiaya(c *gin.Context) {
	PerkiraanBiayaHelper(c, "ULP")
}

func PPEPerkiraanBiaya(c *gin.Context) {
	PerkiraanBiayaHelper(c, "PPE")
}

func PerkiraanBiayaHelper(c *gin.Context, role string) {
	//validate entity must be kelb
	entity, err := services.ValidateTokenFromHeader(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, err.Error())
		return
	}

	if entity.Role != "PPK" {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, `NOT PPK`)
		return
	}

	//get ppk_rp_id and rp_id from params
	val := c.Params.ByName("ppk_rp_id")
	ppkrpid, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		services.SendBasicResponse(c, http.StatusBadRequest, false, err.Error())
		return
	}
	val2 := c.Params.ByName("rp_id")
	rpid, err := strconv.ParseInt(val2, 10, 64)
	if err != nil {
		services.SendBasicResponse(c, http.StatusBadRequest, false, err.Error())
		return
	}

	//decode payload
	estCostString := c.PostForm("est_cost")

	estCostFloat, err := strconv.ParseFloat(estCostString, 64)
	if err != nil {
		services.SendBasicResponse(c, http.StatusBadRequest, false, err.Error())
		return
	}

	doc, err := c.FormFile("doc")
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	docPath := fmt.Sprintf("archive/perkiraan-biaya/%s%s", entity.Fullname, time.Now())

	err = c.SaveUploadedFile(doc, docPath)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	perkiraanBiaya := models.PerkiraanBiaya{
		CreatorID: entity.ID,
		RPID:      rpid,
		EstCost:   estCostFloat,
		Doc:       docPath,
	}

	//validate not 0
	if perkiraanBiaya.EstCost <= 0 {
		services.RemoveFile(docPath)
		services.SendBasicResponse(c, http.StatusInternalServerError, false, "est_cost <= 0")
		return
	}

	//extract db
	db, err := services.GetDB(c)
	if err != nil {
		services.RemoveFile(docPath)
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//store perkiraan biaya
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		services.RemoveFile(docPath)
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}
	res, err := perkiraanBiaya.InsertTx(tx, ctx, entity.ID)
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

	//store ulp_perkiraan_biaya or ppe_perkiraan_biaya
	switch role {
	case "ULP":
		ulpperkiraanbiaya := models.ULPPerkiraanBiaya{
			PerkiraanBiayaID: lastInsertedID,
		}
		_, err = ulpperkiraanbiaya.InsertTx(tx, ctx)
	case "PPE":
		ppeperkiraanbiaya := models.PPEPerkiraanBiaya{
			PerkiraanBiayaID: lastInsertedID,
		}
		_, err = ppeperkiraanbiaya.InsertTx(tx, ctx)
	}
	if err != nil {
		services.RemoveFile(docPath)
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//change status of rp
	rp := models.RP{
		Status: "ON ROUTE TO PPE / ULP",
	}
	_, err = rp.UpdateStatusTx(tx, ctx)
	if err != nil {
		services.RemoveFile(docPath)
		tx.Rollback()
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//delete ppk_rp
	ppkrp := models.PPKRP{
		ID: ppkrpid,
	}
	_, err = ppkrp.DeleteTx(tx, ctx)
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
	services.SendBasicResponse(c, http.StatusOK, true, "perkiraan biaya is created and routed to ppe/ulp")
}
