package services

import (
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
)

func GetDB(c *gin.Context) (*sql.DB, error) {
	val, exist := c.Get("DB")
	if !exist {
		return nil, errors.New("NO DB IN MIDDLEWARE")
	}

	return val.(*sql.DB), nil
}
