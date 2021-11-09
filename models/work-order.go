package models

import "time"

type WorkOrder struct {
	ID           int64 `json:"id"`
	WorkRequest  WorkRequest
	EstDate      time.Time `json:"est_date"`
	EstLaborHour int64     `json:"est_labor_hour"`
	Worker       string    `json:"worker"`
	Worker_Email string    `json:"worker_email"`
	Cost         float64   `json:"cost"`
	Status       int64     `json:"status"`
	ApproverID   int64     `json:"approver_id"`
}

func (w *WorkOrder) Create() {

}

func (w *WorkOrder) Close() {

}

func (w *WorkOrder) Delete() {

}
