package models

type PUMInbox struct {
	ID        int64 `json:"pum_inbox_id"`
	WorkOrder WorkOrder
}

func (w *PUMInbox) Create() {

}
