package controllers

import (
	"context"
	"encoding/json"
	"manajemen-work-order/models"
	"manajemen-work-order/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChangePWD(c *gin.Context) {
	//validate entity must be bdmu
	entity, err := services.ValidateTokenFromHeader(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, err.Error())
		return
	}

	//decode payload
	payload := models.Entity{}
	err = json.NewDecoder(c.Request.Body).Decode(&payload)
	if err != nil {
		services.SendBasicResponse(c, http.StatusBadRequest, false, err.Error())
		return
	}
	err = services.IsNotEmpty(payload.Password)
	if err != nil {
		services.SendBasicResponse(c, http.StatusBadRequest, false, err.Error())
		return
	}

	//hash the password
	hashedPass, err := services.HashPassword(&payload.Password)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//extract db
	ctx := context.Background()
	db, err := services.GetDB(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//update password in db
	payload.ID = entity.ID
	payload.Password = hashedPass
	err = payload.ChangePWD(db, ctx)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//respond
	services.SendBasicResponse(c, http.StatusOK, true, "password changed")
}
