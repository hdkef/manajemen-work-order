package main

import (
	"manajemen-work-order/controllers"
	"manajemen-work-order/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.LoadHTMLGlob("view/**/*")
	r.Static("/assets", "./assets")
	r.Use(middlewares.Auth)

	r.GET("/login", func(c *gin.Context) {
		controllers.Login(c)
	})
	r.GET("/wo-detail/:id", func(c *gin.Context) {
		controllers.WODetail(c)
	})
	r.GET("/wr-detail/:id", func(c *gin.Context) {
		controllers.WRDetail(c)
	})
	r.GET("/wo-progress/:id", func(c *gin.Context) {
		controllers.WOProgress(c)
	})
	r.GET("/ppe-dashboard", func(c *gin.Context) {
		controllers.PPEDashboard(c)
	})
	r.GET("/ppk-dashboard", func(c *gin.Context) {
		controllers.PPKDashboard(c)
	})
	r.GET("/pum-dashboard", func(c *gin.Context) {
		controllers.PUMDashboard(c)
	})
	r.GET("/user-dashboard", func(c *gin.Context) {
		controllers.UserDashboard(c)
	})

	r.POST("/login", func(c *gin.Context) {
		controllers.Login(c)
	})

	r.Run()
}
