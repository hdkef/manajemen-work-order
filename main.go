package main

import (
	"manajemen-work-order/controllers"
	"manajemen-work-order/services"

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

	ulpperkiraanbiaya := ulp.Group("/perkiraan_biaya")
	ppeperkiraanbiaya := ppe.Group("/perkiraan_biaya")

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

	rp.POST("/new/:ppp_id/:kelb_ppp_id", controllers.RP)

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
	kela.DELETE("/:id", controllers.KELAPPPDelete)
	kelb.DELETE("/:id", controllers.KELBPPPDelete)
	bdmurp.DELETE("/:id", controllers.BDMURPDelete)
	bdmuprp.DELETE("/:id", controllers.BDMUPRPDelete)
	kela.DELETE("/:id", controllers.KELARPDelete)

	//route edit

	spk.PUT("/:id", controllers.SPKEdit)

	//route get
	bdmuppp.GET("", controllers.BDMUPPPGet)
	bdmupppp.GET("", controllers.BDMUPPPPGet)
	kelappp.GET("", controllers.KELAPPPGet)
	kelbppp.GET("", controllers.KELBPPPGet)
	bdmurp.GET("", controllers.BDMURPGet)
	bdmuprp.GET("", controllers.BDMUPRPGet)
	kelarp.GET("", controllers.KELARPGet)
	ulpperkiraanbiaya.GET("", controllers.ULPPekiraanBiayaGet)
	ppeperkiraanbiaya.GET("", controllers.PPEPekiraanBiayaGet)
	ppkpengadaan.GET("", controllers.PPKPengadaanGet)
	ppkspk.GET("", controllers.PPKSPKGet)

	r.Run()
}
