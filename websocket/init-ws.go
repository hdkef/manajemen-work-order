package websocket

import (
	"context"
	"fmt"
	"log"
	"manajemen-work-order/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

//const for pingponger

const (
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

//onlineMap to store websocket.Conn with ID key

var onlineMap map[int64]*websocket.Conn = make(map[int64]*websocket.Conn)

//various channel to handle various payload type
var initUserFromClientChan chan models.Message = make(chan models.Message)

//upgrader to upgrade http to websocket.Conn
var upgrader websocket.Upgrader = websocket.Upgrader{
	CheckOrigin: func(req *http.Request) bool {
		return true
	},
}

//only upgrade and initiate websocket if there is user context in middleware
func InitWS(c *gin.Context) {

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	//create goroutine as reader and sender payload
	go readAndSend(cancel, ws)

	//create goroutine as receiver and handler
	go receiveAndHandle(ctx)
}

//readAndSend read incoming payload, assign websocket.Conn to payload and send to corresponding channel
func readAndSend(cancel context.CancelFunc, ws *websocket.Conn) {
	var payload models.Message = models.Message{
		Conn: ws,
	}
	defer cancel()

	for {
		err := ws.ReadJSON(&payload)
		if err != nil {
			log.Println(err)
			break
		}
		switch payload.Type {
		case "initUserFromClient":
			initUserFromClientChan <- payload
		}
	}
}

//receiveAndHandle receive payload from channel and handle to corresponding function
func receiveAndHandle(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-initUserFromClientChan:
			initUserFromClient(msg)
		}
	}
}

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
