package controllers

import (
	"context"
	"encoding/json"
	"manajemen-work-order/models"
	"manajemen-work-order/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	//decode JSON
	var entity models.Entity
	err := json.NewDecoder(c.Request.Body).Decode(&entity)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//get payload pass
	payloadPass := entity.Password

	//validate data
	err = services.IsNotEmpty(payloadPass, entity.Username)
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

	//find data in db
	ctx := context.Background()
	err = entity.FindOne(db, ctx, "username", entity.Username)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//compare password from payload and database
	err = services.CompareTwoPassword(&payloadPass, &entity.Password)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//create token
	token, err := services.GenerateTokenFromEntity(&entity)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//set token as cookie
	c.SetCookie("Authorization", token, 100000, "/", "", false, false)

	//send
	services.SendBasicResponse(c, http.StatusOK, true, "login success")
}
