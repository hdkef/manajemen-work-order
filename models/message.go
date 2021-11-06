package models

import "github.com/gorilla/websocket"

type Message struct {
	Type  string `json:"type"`
	OK    bool   `json:"ok"`
	Msg   string `json:"msg"`
	Token string `json:"token"`
	Data  interface{}
	Conn  *websocket.Conn
	User  User
}
