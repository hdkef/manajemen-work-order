package websocket

import (
	"manajemen-work-order/mockup"
	"manajemen-work-order/models"
	"manajemen-work-order/utils"
)

//initUserFromClient handle initUserFromClient payload type to be responded with work request that been created by user
func initUserFromClient(payload models.Message) {

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

	utils.WSResponse(payload, "initUserFromServer", true, "", []models.WorkRequest{mockup.WORK_REQUEST, mockup.WORK_REQUEST, mockup.WORK_REQUEST})
}
