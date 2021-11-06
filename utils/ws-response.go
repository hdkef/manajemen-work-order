package utils

import (
	"manajemen-work-order/models"
)

func WSResponse(payload models.Message, resType string, OK bool, Msg string, Data interface{}) {
	response := models.Message{
		Type: resType,
		OK:   OK,
		Msg:  Msg,
		Data: Data,
	}

	payload.Conn.WriteJSON(&response)
}
