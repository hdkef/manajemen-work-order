package websocket

import (
	"context"
	"log"
	"manajemen-work-order/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

//various channel to handle various payload type
var initUserFromClientChan chan models.Message = make(chan models.Message)
var pagingUserFromClientChan chan models.Message = make(chan models.Message)
var createWRUserFromClientChan chan models.Message = make(chan models.Message)
var cancelWRUserFromClientChan chan models.Message = make(chan models.Message)

//only upgrade and initiate websocket
func InitWSUser(c *gin.Context) {

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	//create goroutine as reader and sender payload
	go readAndSendUser(cancel, ws)

	//create goroutine as receiver and handler
	go receiverAndHandleUser(ctx)
}

//readAndSendUser read incoming payload, assign websocket.Conn to payload and send to corresponding channel
func readAndSendUser(cancel context.CancelFunc, ws *websocket.Conn) {
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
		case "pagingUserFromClient":
			pagingUserFromClientChan <- payload
		case "createWRUserFromClient":
			createWRUserFromClientChan <- payload
		case "cancelWRUserFromClient":
			cancelWRUserFromClientChan <- payload
		case "changePasswordFromClient":
			changePasswordFromClientChan <- payload
		}
	}
}

//receiverAndHandleUser receive payload from channel and handle to corresponding function
func receiverAndHandleUser(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-initUserFromClientChan:
			initUserFromClient(msg)
		case msg := <-pagingUserFromClientChan:
			pagingUserFromClient(msg)
		case msg := <-createWRUserFromClientChan:
			createWRUserFromClient(msg)
		case msg := <-cancelWRUserFromClientChan:
			cancelWRUserFromClient(msg)
		case msg := <-changePasswordFromClientChan:
			changePasswordFromClient(msg)
		}
	}
}
