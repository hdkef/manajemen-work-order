package models

type PPEInbox struct {
	ID          int64   `json:"ppe_inbox_id"`
	EstCost     float64 `json:"ppe_inbox_est_cost"`
	WorkRequest WorkRequest
}

func (w *PPEInbox) Create() {

}
