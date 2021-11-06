package models

import "time"

type WorkRequest struct {
	ID          int64     `json:"work_request_id"`
	RequestorID int64     `json:"work_request_requestor"`
	Priority    int64     `json:"work_request_priority"`
	DateCreated time.Time `json:"work_request_date_created"`
	Task        string    `json:"work_request_task"`
	Equipment   string    `json:"work_request_equipment"`
	Location    string    `json:"work_request_location"`
	Instruction string    `json:"work_request_instruction"`
	Description string    `json:"work_request_description"`
	Status      int64     `json:"work_request_status"`
}

func (w *WorkRequest) Create() {

}

func (w *WorkRequest) ChangeStatus() {

}

func (w *WorkRequest) Delete() {

}
