package models

import (
	"time"

	"github.com/gorilla/websocket"
)

type Message struct {
	Type             string           `json:"type"`
	OK               bool             `json:"ok"`
	Msg              string           `json:"msg"`
	Token            string           `json:"token"`
	LastID           int64            `json:"last_id"`
	ChangePWD        ChangePWD        `json:"changepwdfromclient"`
	WRFromClient     WorkRequest      `json:"wrfromclient"`
	WOFromClient     WorkOrder        `json:"wofromclient"`
	UserRespond      UserRespond      `json:"userrespondfromclient"`
	PUMRespond       PUMRespond       `json:"pumrespondfromclient"`
	PPERespond       PPERespond       `json:"pperespondfromclient"`
	PPKRespondPUM    PPKRespondPUM    `json:"ppkrespondpumfromclient"`
	PPKRespondPPE    PPKRespondPPE    `json:"ppkrespondppefromclient"`
	PPKRespondWorker PPKRespondWorker `json:"ppkrespondworkerfromclient"`
	Data             interface{}      `json:"data"`
	Conn             *websocket.Conn
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

type PPKRespondPUM struct {
	ID      int64   `json:"id"`
	EstCost float64 `json:"est_cost"`
}

type PPKRespondPPE struct {
	ID int64 `json:"id"`
}

type PPKRespondWorker struct {
	ID          int64  `json:"id"`
	WorkOrderID int64  `json:"work_order_id"`
	Msg         string `json:"msg"`
	Type        bool   `json:"type"`
}
