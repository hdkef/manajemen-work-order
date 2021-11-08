package websocket

import (
	"context"
	"log"
	"manajemen-work-order/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

//various channel to handle various payload type
var initPUMFromClientChan chan models.Message = make(chan models.Message)
var acceptWRPUMFromClientChan chan models.Message = make(chan models.Message)
var declineWRPUMFromClientChan chan models.Message = make(chan models.Message)
var loadHistoryPUMFromClientChan chan models.Message = make(chan models.Message)

//only upgrade and initiate websocket
func InitWSPUM(c *gin.Context) {

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	//create goroutine as reader and sender payload
	go readAndSendPUM(cancel, ws)

	//create goroutine as receiver and handler
	go receiverAndHandlePUM(ctx)
}

//readAndSendPUM read incoming payload, assign websocket.Conn to payload and send to corresponding channel
func readAndSendPUM(cancel context.CancelFunc, ws *websocket.Conn) {
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
		case "initPUMFromClient":
			initPUMFromClientChan <- payload
		case "acceptWRPUMFromClient":
			acceptWRPUMFromClientChan <- payload
		case "declineWRPUMFromClient":
			declineWRPUMFromClientChan <- payload
		case "changePasswordFromClient":
			changePasswordFromClientChan <- payload
		case "loadHistoryPUMFromClient":
			loadHistoryPUMFromClientChan <- payload
		}
	}
}

//receiverAndHandlePUM receive payload from channel and handle to corresponding function
func receiverAndHandlePUM(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-initPUMFromClientChan:
			initPUMFromClient(msg)
		case msg := <-acceptWRPUMFromClientChan:
			acceptWRPUMFromClient(msg)
		case msg := <-declineWRPUMFromClientChan:
			declineWRPUMFromClient(msg)
		case msg := <-changePasswordFromClientChan:
			changePasswordFromClient(msg)
		case msg := <-loadHistoryPUMFromClientChan:
			loadHistoryPUMFromClient(msg)
		}
	}
}
