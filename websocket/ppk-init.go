package websocket

import (
	"context"
	"log"
	"manajemen-work-order/models"
	"manajemen-work-order/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

//various channel to handle various payload type
var initPPKFromClientChan chan models.Message = make(chan models.Message)
var inboxPUMFromClientChan chan models.Message = make(chan models.Message)
var inboxPPEFromClientChan chan models.Message = make(chan models.Message)
var inboxWorkerFromClientChan chan models.Message = make(chan models.Message)
var ppkRespondPUMFromClientChan chan models.Message = make(chan models.Message)
var createWOFromClientChan chan models.Message = make(chan models.Message)
var ppkRespondWorkerFromClientChan chan models.Message = make(chan models.Message)

//only upgrade and initiate websocket
func InitWSPPK(c *gin.Context) {

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	//create goroutine as reader and sender payload
	go readAndSendPPK(cancel, ws)

	//create goroutine as receiver and handler
	go receiverAndHandlePPK(ctx)
}

//readAndSendPPK read incoming payload, assign websocket.Conn to payload and send to corresponding channel
func readAndSendPPK(cancel context.CancelFunc, ws *websocket.Conn) {
	var payload models.Message = models.Message{
		Conn: ws,
	}
	defer cancel()

	for {
		err := ws.ReadJSON(&payload)
		if err != nil {
			utils.WSResponse(payload, "error", false, err.Error(), nil)
			log.Println(err)
			break
		}

		switch payload.Type {
		case "initPPKFromClient":
			initPPKFromClientChan <- payload
		case "inboxPUMFromClient":
			inboxPUMFromClientChan <- payload
		case "inboxPPEFromClient":
			inboxPPEFromClientChan <- payload
		case "inboxWorkerFromClient":
			inboxWorkerFromClientChan <- payload
		case "ppkRespondPUMFromClient":
			ppkRespondPUMFromClientChan <- payload
		case "createWOFromClient":
			createWOFromClientChan <- payload
		case "ppkRespondWorkerFromClient":
			ppkRespondWorkerFromClientChan <- payload
		}
	}
}

//receiverAndHandlePPK receive payload from channel and handle to corresponding function
func receiverAndHandlePPK(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-initPPKFromClientChan:
			initPPKFromClient(msg)
		case msg := <-inboxPUMFromClientChan:
			inboxPUMFromClient(msg)
		case msg := <-inboxPPEFromClientChan:
			inboxPPEFromClient(msg)
		case msg := <-inboxWorkerFromClientChan:
			inboxWorkerFromClient(msg)
		case msg := <-ppkRespondPUMFromClientChan:
			ppkRespondPUMFromClient(msg)
		case msg := <-createWOFromClientChan:
			createWOFromClient(msg)
		case msg := <-ppkRespondWorkerFromClientChan:
			ppkRespondWorkerFromClient(msg)
		}
	}
}
