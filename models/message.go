package models

import (
	"time"

	"github.com/gorilla/websocket"
)

type Message struct {
	Type         string      `json:"type"`
	OK           bool        `json:"ok"`
	Msg          string      `json:"msg"`
	Token        string      `json:"token"`
	LastID       int64       `json:"last_id"`
	ChangePWD    ChangePWD   `json:"changepwdfromclient"`
	WRFromClient WorkRequest `json:"wrfromclient"`
	UserRespond  UserRespond `json:"userrespondfromclient"`
	PUMRespond   PUMRespond  `json:"pumrespondfromclient"`
	PPERespond   PPERespond  `json:"pperespondfromclient"`
	Data         interface{} `json:"data"`
	Conn         *websocket.Conn
	User         User
}

type ChangePWD struct {
	OldPassword string `json:"old"`
	NewPassword string `json:"new"`
}

type UserRespond struct {
	ID int64 `json:"id"`
}

type PUMRespond struct {
	ID int64 `json:"id"`
}

type PPERespond struct {
	ID           int64     `json:"id"`
	EstDate      time.Time `json:"est_date"`
	EstLaborHour int64     `json:"est_labor_hour"`
	Worker       string    `json:"worker"`
	WorkerEmail  string    `json:"worker_email"`
	Cost         float64   `json:"cost"`
}
