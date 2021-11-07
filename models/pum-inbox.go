package models

type PUMInbox struct {
	ID          int64 `json:"pum_inbox_id"`
	WorkRequest WorkRequest
}

func (w *PUMInbox) Create() {

}
