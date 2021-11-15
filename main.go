package main

import (
	"manajemen-work-order/controllers"
	"manajemen-work-order/services"
	"manajemen-work-order/views"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := services.DB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := gin.New()

	//add DB to gin context middleware
	r.Use(func(c *gin.Context) {
		c.Set("DB", db)
	})

	r.Static("archive", "archive")
	r.Static("assets/signature", "assets/signature")
	r.Static("assets/img", "assets/img")
	r.Static("/js", "./assets/js")
	r.Static("/css", "./assets/css")
	r.Static("/img", "./assets/img")
	r.LoadHTMLGlob("views/*")

	//backend
	api := r.Group("/api/v1")
	entity := api.Group("/entity")
	ppp := api.Group("/ppp")
	rp := api.Group("/rp")
	perkiraanBiaya := api.Group("/perkiraan-biaya")
	spk := api.Group("/spk")
	pengadaan := api.Group("/pengadaan")
	bdmu := api.Group("/bdmu")
	bdmup := api.Group("/bdmup")
	kela := api.Group("/kela")
	kelb := api.Group("/kelb")
	ppk := api.Group("/ppk")
	ppe := api.Group("/ppe")
	ulp := api.Group("/ulp")

	bdmuppp := bdmu.Group("/ppp")
	bdmupppp := bdmup.Group("/ppp")
	kelappp := kela.Group("/ppp")
	kelbppp := kelb.Group("/ppp")

	bdmurp := bdmu.Group("/rp")
	bdmuprp := bdmup.Group("/rp")
	kelarp := kela.Group("/rp")
	ppkrp := ppk.Group("/rp")

	ulpperkiraanbiaya := ulp.Group("/perkiraan-biaya")
	ppeperkiraanbiaya := ppe.Group("/perkiraan-biaya")

	ppkpengadaan := ppk.Group("/pengadaan")
	ppkspk := ppk.Group("/spk")

	//post route
	api.POST("/login", controllers.Login)

	entity.POST("", controllers.EntityPost)

	ppp.POST("", controllers.PPPPost)
	ppp.POST("/:ppp_id/ok/bdmu/:bdmu_id", controllers.PPPOKBDMU)
	ppp.POST("/:ppp_id/ok/bdmup/:bdmup_id", controllers.PPPOKBDMUP)
	ppp.POST("/:ppp_id/ok/kela/:kela_id", controllers.PPPOKKELA)

	ppp.POST("/:ppp_id/no", controllers.PPPNO)

	rp.POST("/new/:ppp_id/:kelb_ppp_id", controllers.RPPost)

	rp.POST("/:rp_id/ok/kela/:kela_id", controllers.RPOKKELA)
	rp.POST("/:rp_id/ok/bdmup/:bdmup_id", controllers.RPOKBDMUP)
	rp.POST("/:rp_id/ok/bdmu/:bdmu_id", controllers.RPOKBDMU)

	rp.POST("/:rp_id/no", controllers.RPNO)

	perkiraanBiaya.POST("/ulp/:rp_id/:ppk_rp_id", controllers.ULPPerkiraanBiaya)
	perkiraanBiaya.POST("/ppe/:rp_id/:ppk_rp_id", controllers.PPEPerkiraanBiaya)

	spk.POST("/new/:pengadaan_id/:ppk_pengadaan_id", controllers.SPKPost)
	spk.POST("/:id/lapor", controllers.SPKLapor)
	spk.POST("/:id/ok/:ppk_spk_id", controllers.SPKOK)
	spk.POST("/:id/no/:ppk_spk_id", controllers.SPKNO)

	pengadaan.POST("/:perkiraan_biaya_id/ulp/:inbox_id", controllers.PengadaanFromULP)
	pengadaan.POST("/:perkiraan_biaya_id/ppe/:inbox_id", controllers.PengadaanFromPPE)
	//route delete
	entity.DELETE("/:id", controllers.EntityDelete)
	bdmuppp.DELETE("/:id", controllers.BDMUPPPDelete)
	bdmupppp.DELETE("/:id", controllers.BDMUPPPPDelete)
	kelappp.DELETE("/:id", controllers.KELAPPPDelete)
	kelbppp.DELETE("/:id", controllers.KELBPPPDelete)
	bdmurp.DELETE("/:id", controllers.BDMURPDelete)
	bdmuprp.DELETE("/:id", controllers.BDMUPRPDelete)
	kelarp.DELETE("/:id", controllers.KELARPDelete)
	ppkrp.DELETE("/:id", controllers.PPKRPDelete)

	//route edit

	spk.PUT("/:id", controllers.SPKEdit)
	api.PUT("/changepwd", controllers.ChangePWD)

	//route get
	entity.GET("", controllers.EntityGet)

	bdmuppp.GET("", controllers.BDMUPPPGet)
	bdmupppp.GET("", controllers.BDMUPPPPGet)
	kelappp.GET("", controllers.KELAPPPGet)
	kelbppp.GET("", controllers.KELBPPPGet)
	bdmurp.GET("", controllers.BDMURPGet)
	bdmuprp.GET("", controllers.BDMUPRPGet)
	ppkrp.GET("", controllers.PPKRPGet)
	kelarp.GET("", controllers.KELARPGet)
	ulpperkiraanbiaya.GET("", controllers.ULPPekiraanBiayaGet)
	ppeperkiraanbiaya.GET("", controllers.PPEPekiraanBiayaGet)
	ppkpengadaan.GET("", controllers.PPKPengadaanGet)
	ppkspk.GET("", controllers.PPKSPKGet)

	ppp.GET("", controllers.PPPGet)
	rp.GET("", controllers.RPGet)
	perkiraanBiaya.GET("", controllers.PerkiraanBiayaGet)
	pengadaan.GET("", controllers.PengadaanGet)
	spk.GET("", controllers.SPKGet)

	//frontend
	r.GET("", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
	})
	r.GET("/login", views.Login)
	r.GET("/create-entity", views.CreateEntity)
	r.GET("/create-ppp", views.CreatePPP)
	r.GET("/create-rp/:id/:ppp_id", views.CreateRP)
	r.GET("/create-spk/:id/:pengadaan_id", views.CreateSPK)
	r.GET("/create-perkiraan-biaya/:id/:rp_id", views.CreatePerkiraanBiaya)
	r.GET("/create-pengadaan/:id/:perkiraan_biaya_id", views.CreatePengadaan)

	r.GET("/change-pwd", views.ChangePWD)
	r.GET("/spk-progress/:id", views.SPKProgress)
	r.GET("/revision/:id/:spk_id", views.Revision)

	r.GET("/user", views.User)
	r.GET("/bdmu", views.BDMU)
	r.GET("/bdmup", views.BDMUP)
	r.GET("/kela", views.KELA)
	r.GET("/kelb", views.KELB)
	r.GET("/ppk", views.PPK)
	r.GET("/ppe", views.PPE)
	r.GET("/ulp", views.ULP)
	r.GET("/super-admin", views.SuperAdmin)

	r.GET("/ppp/:id", views.PPPOne)
	r.GET("/rp/:id", views.RPOne)
	r.GET("/perkiraan-biaya/:id", views.PerkiraanBiayaOne)
	r.GET("/pengadaan/:id", views.PengadaanOne)
	r.GET("/spk/:id", views.SPKOne)
	r.GET("/entity/:id", views.EntityOne)

	r.GET("/ppp", views.PPPAll)
	r.GET("/rp", views.RPAll)
	r.GET("/perkiraan-biaya", views.PerkiraanBiayaAll)
	r.GET("/pengadaan", views.PengadaanAll)
	r.GET("/spk", views.SPKAll)

	r.Run()
}
