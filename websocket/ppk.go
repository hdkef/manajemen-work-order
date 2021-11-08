package websocket

import (
	"manajemen-work-order/mockup"
	"manajemen-work-order/models"
	"manajemen-work-order/utils"
)

func initPPKFromClient(payload models.Message) {
	//auth
	user := models.User{}

	err := user.ValidateTokenStringGetUser(&payload.Token)
	if err != nil {
		utils.WSResponse(payload, "error", false, "unauthorized", nil)
		payload.Conn.Close()
		return
	}

	//assign id websocket conn to online map
	onlineMap[payload.User.ID] = payload.Conn

	//create goroutine for ping ponger
	go pingPonger(payload.User.ID, payload.Conn)
}

func inboxPUMFromClient(payload models.Message) {
	//auth
	user := models.User{}

	err := user.ValidateTokenStringGetUser(&payload.Token)
	if err != nil {
		utils.WSResponse(payload, "error", false, "unauthorized", nil)
		payload.Conn.Close()
		return
	}
	utils.WSResponse(payload, "inboxPUMFromServer", true, "", mockup.PPKInboxFromPUMSlice(5, 1))
}

func inboxPPEFromClient(payload models.Message) {
	//auth
	user := models.User{}

	err := user.ValidateTokenStringGetUser(&payload.Token)
	if err != nil {
		utils.WSResponse(payload, "error", false, "unauthorized", nil)
		payload.Conn.Close()
		return
	}
	utils.WSResponse(payload, "inboxPPEFromServer", true, "", mockup.PPKInboxFromPPESlice(5, 1))
}

func inboxWorkerFromClient(payload models.Message) {
	//auth
	user := models.User{}

	err := user.ValidateTokenStringGetUser(&payload.Token)
	if err != nil {
		utils.WSResponse(payload, "error", false, "unauthorized", nil)
		payload.Conn.Close()
		return
	}
	utils.WSResponse(payload, "inboxWorkerFromServer", true, "", mockup.PPKInboxFromWorkerSlice(5, 1))
}
