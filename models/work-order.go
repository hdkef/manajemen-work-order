package models

import "time"

type WorkOrder struct {
	ID           int64 `json:"work_order_id"`
	WorkRequest  WorkRequest
	EstDate      time.Time `json:"work_order_est_date"`
	EstLaborHour int64     `json:"work_order_est_labor_hour"`
	Worker       string    `json:"work_order_worker"`
	Worker_Email string    `json:"work_order_worker_email"`
	Cost         float64   `json:"work_order_cost"`
	Status       string    `json:"work_order_status"`
	ApproverID   int64     `json:"work_order_approver_id"`
}

func (w *WorkOrder) Create() {

}

func (w *WorkOrder) Close() {

}

func (w *WorkOrder) Delete() {

}
