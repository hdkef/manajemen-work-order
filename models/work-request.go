package models

import "time"

type WorkRequest struct {
	ID          int64     `json:"id"`
	RequestorID int64     `json:"requestor"`
	Priority    int64     `json:"priority"`
	DateCreated time.Time `json:"date_created"`
	Task        string    `json:"task"`
	Equipment   string    `json:"equipment"`
	Location    string    `json:"location"`
	Instruction string    `json:"instruction"`
	Description string    `json:"description"`
	Status      int64     `json:"status"`
}

func (w *WorkRequest) Create() {

}

func (w *WorkRequest) ChangeStatus() {

}

func (w *WorkRequest) Delete() {

}
