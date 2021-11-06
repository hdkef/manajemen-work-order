package models

type PPEInbox struct {
	ID        int64 `json:"ppe_inbox_id"`
	WorkOrder WorkOrder
}
