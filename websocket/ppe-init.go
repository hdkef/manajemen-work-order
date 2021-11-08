package websocket

import (
	"context"
	"log"
	"manajemen-work-order/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

//various channel to handle various payload type
var initPPEFromClientChan chan models.Message = make(chan models.Message)

//only upgrade and initiate websocket
func InitWSPPE(c *gin.Context) {

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	//create goroutine as reader and sender payload
	go readAndSendPPE(cancel, ws)

	//create goroutine as receiver and handler
	go receiverAndHandlePPE(ctx)
}

//readAndSendPPE read incoming payload, assign websocket.Conn to payload and send to corresponding channel
func readAndSendPPE(cancel context.CancelFunc, ws *websocket.Conn) {
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
		case "initPPEFromClient":
			initPPEFromClientChan <- payload
		}
	}
}

//receiverAndHandlePPE receive payload from channel and handle to corresponding function
func receiverAndHandlePPE(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-initPPEFromClientChan:
			initPPEFromClient(msg)
		}
	}
}
