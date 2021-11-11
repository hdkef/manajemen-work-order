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

	//post route
	api.POST("/login", controllers.Login)

	entity.POST("", controllers.EntityPost)

	ppp.POST("", controllers.PPPPost)
	ppp.POST("/:ppp_id/ok/bdmu/:bdmu_id", controllers.PPPOKBDMU)
	ppp.POST("/:ppp_id/no/bdmu/:bdmu_id", controllers.PPPNOBDMU)
	ppp.POST("/:ppp_id/ok/bdmup/:bdmup_id", controllers.PPPOKBDMUP)
	ppp.POST("/:ppp_id/no/bdmup/:bdmup_id", controllers.PPPNOBDMUP)
	ppp.POST("/:ppp_id/ok/kela/:kela_id", controllers.PPPOKKELA)
	ppp.POST("/:ppp_id/no/kela/:kela_id", controllers.PPPNOKELA)

	rp.POST("/new/:ppp_id/:kelb_ppp_id", controllers.RP)

	rp.POST("/:rp_id/ok/kela/:kela_id", controllers.RPOKKELA)
	rp.POST("/:rp_id/no/kela/:kela_id", controllers.RPNOKELA)
	rp.POST("/:rp_id/ok/bdmup/:bdmup_id", controllers.RPOKBDMUP)
	rp.POST("/:rp_id/no/bdmup/:bdmup_id", controllers.RPNOBDMUP)
	rp.POST("/:rp_id/ok/bdmu/:bdmu_id", controllers.RPOKBDMU)
	rp.POST("/:rp_id/no/bdmu/:bdmu_id", controllers.RPNOBDMU)

	perkiraanBiaya.POST("/ulp/:rp_id/:ppk_rp_id", controllers.ULPPerkiraanBiaya)
	perkiraanBiaya.POST("/ppe/:rp_id/:ppk_rp_id", controllers.PPEPerkiraanBiaya)

	spk.POST("/new/:pengadaan_id/:ppk_pengadaan_id", controllers.SPKPost)
	spk.POST("/:id/lapor", controllers.SPKLapor)
	spk.POST("/:id/ok", controllers.SPKOK)
	spk.POST("/:id/no", controllers.SPKNO)

	pengadaan.POST("/:perkiraan_biaya_id/ulp/:inbox_id", controllers.PengadaanFromULP)
	pengadaan.POST("/:perkiraan_biaya_id/ppe/:inbox_id", controllers.PengadaanFromPPE)

	//route delete
	entity.DELETE("/:id", controllers.EntityDelete)

	//route edit

	spk.PUT("/:id", controllers.SPKEdit)

	r.Run()
}
