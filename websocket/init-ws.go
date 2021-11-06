package websocket

import (
	"context"
	"log"
	"manajemen-work-order/models"
	"manajemen-work-order/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

//upgrader to upgrade http to websocket.Conn
var upgrader websocket.Upgrader = websocket.Upgrader{
	CheckOrigin: func(req *http.Request) bool {
		return true
	},
}

//only upgrade and initiate websocket if there is user context in middleware
func InitWS(c *gin.Context) {
	val, exist := c.Get("User")
	if !exist {
		utils.Response(c, http.StatusUnauthorized, false, "unauthenticated")
		return
	}

	user := val.(models.User)

	ws, err := upgrader.Upgrade(c.Writer, c.Request, c.Request.Header)
	if err != nil {
		log.Println(err)
		utils.Response(c, http.StatusInternalServerError, false, err.Error())
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	go readAndSend(cancel, ws, &user)
	go receiveAndHandle(ctx)
}

//readAndSend read incoming payload, assign websocket.Conn to payload and send to corresponding channel
func readAndSend(cancel context.CancelFunc, ws *websocket.Conn, user *models.User) {

}

//receiveAndHandle receive payload from channel and handle to corresponding function
func receiveAndHandle(ctx context.Context) {

}
