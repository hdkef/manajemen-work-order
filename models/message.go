package models

import "github.com/gorilla/websocket"

type Message struct {
	Type         string      `json:"type"`
	OK           bool        `json:"ok"`
	Msg          string      `json:"msg"`
	Token        string      `json:"token"`
	LastID       int64       `json:"last_id"`
	WRFromClient WorkRequest `json:"wrfromclient"`
	Data         interface{} `json:"data"`
	Conn         *websocket.Conn
	User         User
}
