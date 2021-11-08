package websocket

import (
	"fmt"
	"manajemen-work-order/models"
	"manajemen-work-order/utils"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

//const for pingponger

const (
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

//onlineMap to store websocket.Conn with ID key

var onlineMap map[int64]*websocket.Conn = make(map[int64]*websocket.Conn)

//upgrader to upgrade http to websocket.Conn
var upgrader websocket.Upgrader = websocket.Upgrader{
	CheckOrigin: func(req *http.Request) bool {
		return true
	},
}

//common channel
var changePasswordFromClientChan chan models.Message = make(chan models.Message)

//pingPonger will ping websocket conn and delete onlineMap if return error for defined time range
func pingPonger(ID int64, ws *websocket.Conn) {

	ws.SetPongHandler(func(appData string) error {
		ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	timer := time.NewTicker(pingPeriod)

	defer func() {
		timer.Stop()
		if onlineMap[ID] == ws {
			fmt.Println("delete")
			delete(onlineMap, ID)
		}
	}()

	for {
		select {
		case <-timer.C:
			fmt.Println("tick")
			if err := ws.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func changePasswordFromClient(payload models.Message) {
	//auth
	user := models.User{}

	err := user.ValidateTokenStringGetUser(&payload.Token)
	if err != nil {
		utils.WSResponse(payload, "error", false, "unauthorized", nil)
		payload.Conn.Close()
		return
	}

	//TOBE
	utils.WSResponse(payload, "changePasswordFromServer", true, "password telah diganti", nil)
}
