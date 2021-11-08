package websocket

import (
	"manajemen-work-order/mockup"
	"manajemen-work-order/models"
	"manajemen-work-order/utils"
)

//initPUMFromClient handle initPUMFromClient payload type to be responded with work request that been created by user
func initPUMFromClient(payload models.Message) {

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

	utils.WSResponse(payload, "initPUMFromServer", true, "", mockup.PUMInboxSlice(5, 1))
}

//handle accept work request by PUM
func acceptWRPUMFromClient(payload models.Message) {
	//auth
	user := models.User{}

	err := user.ValidateTokenStringGetUser(&payload.Token)
	if err != nil {
		utils.WSResponse(payload, "error", false, "unauthorized", nil)
		payload.Conn.Close()
		return
	}

	//TOBE

	utils.WSResponse(payload, "resWRPUMFromServer", true, "work request telah diteruskan", payload.IDFromClient)
}

func declineWRPUMFromClient(payload models.Message) {
	//auth
	user := models.User{}

	err := user.ValidateTokenStringGetUser(&payload.Token)
	if err != nil {
		utils.WSResponse(payload, "error", false, "unauthorized", nil)
		payload.Conn.Close()
		return
	}

	//TOBE

	utils.WSResponse(payload, "resWRPUMFromServer", true, "work request telah ditolak", payload.IDFromClient)
}

func loadHistoryPUMFromClient(payload models.Message) {
	user := models.User{}

	err := user.ValidateTokenStringGetUser(&payload.Token)
	if err != nil {
		utils.WSResponse(payload, "error", false, "unauthorized", nil)
		payload.Conn.Close()
		return
	}

	//TOBE
	utils.WSResponse(payload, "loadHistoryPUMFromServer", true, "", mockup.PUMInboxSlice(5, 1))
}
