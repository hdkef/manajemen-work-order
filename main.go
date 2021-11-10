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
	ppp.POST("/:id/ok/bdmu", controllers.PPPOKBDMU)
	ppp.POST("/:id/no/bdmu", controllers.PPPNOBDMU)
	ppp.POST("/:id/ok/bdmup", controllers.PPPOKBDMUP)
	ppp.POST("/:id/no/bdmup", controllers.PPPNOBDMUP)
	ppp.POST("/:id/ok/kela", controllers.PPPOKKELA)
	ppp.POST("/:id/no/kela", controllers.PPPNOKELA)

	rp.POST("", controllers.RP)
	rp.POST("/:id/ok/bdmu", controllers.RPOKBDMU)
	rp.POST("/:id/no/bdmu", controllers.RPNOBDMU)
	rp.POST("/:id/ok/bdmup", controllers.RPOKBDMUP)
	rp.POST("/:id/no/bdmup", controllers.RPNOBDMUP)
	rp.POST("/:id/ok/kela", controllers.RPOKKELA)
	rp.POST("/:id/no/kela", controllers.RPNOKELA)

	perkiraanBiaya.POST("/:id/ulp", controllers.PerkiraanBiayaULP)
	perkiraanBiaya.POST("/:id/ppe", controllers.PerkiraanBiayaULP)

	spk.POST("", controllers.SPK)
	spk.POST("/:id/lapor", controllers.SPKLapor)
	spk.POST("/:id/ok", controllers.SPKOK)
	spk.POST("/:id/no", controllers.SPKNO)

	pengadaan.POST("", controllers.Pengadaan)

	//route delete
	entity.DELETE("/:id", controllers.EntityDelete)

	r.Run()
}
