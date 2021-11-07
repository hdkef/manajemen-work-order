package models

import "github.com/gorilla/websocket"

type Message struct {
	Type         string      `json:"type"`
	OK           bool        `json:"ok"`
	Msg          string      `json:"msg"`
	Token        string      `json:"token"`
	LastID       int64       `json:"last_id"`
	ChangePWD    ChangePWD   `json:"changepwdfromclient"`
	WRFromClient WorkRequest `json:"wrfromclient"`
	IDFromClient int64       `json:"idfromclient"`
	Data         interface{} `json:"data"`
	Conn         *websocket.Conn
	User         User
}

type ChangePWD struct {
	OldPassword string `json:"old"`
	NewPassword string `json:"new"`
}
