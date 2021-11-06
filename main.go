package main

import (
	"manajemen-work-order/controllers"
	"manajemen-work-order/middlewares"
	"manajemen-work-order/websocket"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.LoadHTMLGlob("view/**/*")
	r.Static("/assets", "./assets")

	withoutJWT := r.Group("")
	withJWT := r.Group("")

	withJWT.Use(middlewares.Auth)

	withJWT.GET("/before-ws", func(c *gin.Context) {
		websocket.BeforeWS(c)
	})
	withJWT.GET("/login", func(c *gin.Context) {
		controllers.Login(c)
	})
	withJWT.GET("/wo-detail/:id", func(c *gin.Context) {
		controllers.WODetail(c)
	})
	withJWT.GET("/wr-detail/:id", func(c *gin.Context) {
		controllers.WRDetail(c)
	})
	withJWT.GET("/wo-progress/:id", func(c *gin.Context) {
		controllers.WOProgress(c)
	})
	withJWT.GET("/ppe-dashboard", func(c *gin.Context) {
		controllers.PPEDashboard(c)
	})
	withJWT.GET("/ppk-dashboard", func(c *gin.Context) {
		controllers.PPKDashboard(c)
	})
	withJWT.GET("/pum-dashboard", func(c *gin.Context) {
		controllers.PUMDashboard(c)
	})
	withJWT.GET("/user-dashboard", func(c *gin.Context) {
		controllers.UserDashboard(c)
	})

	withJWT.POST("/login", func(c *gin.Context) {
		controllers.Login(c)
	})

	withoutJWT.GET("/websocket", func(c *gin.Context) {
		websocket.InitWS(c)
	})

	r.Run()
}
