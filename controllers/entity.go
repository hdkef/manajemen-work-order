package controllers

import (
	"context"
	"fmt"
	"manajemen-work-order/models"
	"manajemen-work-order/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func EntityGet(c *gin.Context) {
	//validate entity that entity role is super-admin
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, err.Error())
		return
	}

	if entity.Role != "Super-Admin" {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, "NOT Super-Admin")
		return
	}

	//extract db
	ctx := context.Background()
	db, err := services.GetDB(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	mdl := models.Entity{}

	entities, err := mdl.FindAll(db, ctx)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	c.JSON(http.StatusOK, entities)
}

func EntityPost(c *gin.Context) {
	//validate entity that entity role is super-admin
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, err.Error())
		return
	}

	if entity.Role != "Super-Admin" {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, "NOT Super-Admin")
		return
	}

	//decode payload
	fullname := c.PostForm("fullname")
	username := c.PostForm("username")
	password := c.PostForm("password")
	role := c.PostForm("role")
	email := c.PostForm("email")

	//validate for empty
	err = services.IsNotEmpty(fullname, username, password, role, email)
	if err != nil {
		services.SendBasicResponse(c, http.StatusBadRequest, false, err.Error())
		return
	}

	signature, err := c.FormFile("signature")
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//validation for email
	err = services.IsEmail(email)
	if err != nil {
		services.SendBasicResponse(c, http.StatusBadRequest, false, err.Error())
		return
	}

	signaturePath := fmt.Sprintf("assets/signature/%s%s", fullname, signature.Filename)

	err = c.SaveUploadedFile(signature, signaturePath)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	payload := models.Entity{
		Fullname:  fullname,
		Username:  username,
		Password:  password,
		Role:      role,
		Email:     email,
		Signature: signaturePath,
	}

	//hash password
	hashedPass, err := services.HashPassword(&payload.Password)
	if err != nil {
		services.RemoveFile(signaturePath)
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}
	payload.Password = hashedPass

	//extract db
	db, err := services.GetDB(c)
	if err != nil {
		services.RemoveFile(signaturePath)
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//store entity to database
	ctx := context.Background()
	err = payload.Insert(db, ctx)
	if err != nil {
		services.RemoveFile(signaturePath)
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//respond
	services.SendBasicResponse(c, http.StatusOK, true, "entity created")
}

func EntityDelete(c *gin.Context) {
	//validate entity that entity role is super-admin
	entity, err := services.ValidateTokenFromCookie(c)
	if err != nil {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, err.Error())
		return
	}

	if entity.Role != "Super-Admin" {
		services.SendBasicResponse(c, http.StatusUnauthorized, false, "NOT Super-Admin")
		return
	}

	//extract id from path and validate
	val := c.Params.ByName("id")
	err = services.IsNotEmpty(val)
	if err != nil {
		services.SendBasicResponse(c, http.StatusBadRequest, false, err.Error())
		return
	}

	id, err := strconv.ParseInt(val, 10, 64)
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
	entityModel := models.Entity{
		ID: id,
	}

	//get entity
	signature, err := entityModel.FindSignature(db, ctx)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//delete entity from disk
	services.RemoveFile(signature)

	//delete entity
	err = entityModel.Delete(db, ctx)
	if err != nil {
		services.SendBasicResponse(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	//respond
	services.SendBasicResponse(c, http.StatusOK, true, "entity deleted")
}
