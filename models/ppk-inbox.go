package models

import "time"

type PPKInbox struct {
	FromPUM    FromPUM
	FromPPE    FromPPE
	FromWorker FromWorker
}

type FromPUM struct {
	ID          int64     `json:"ppk_inbox_from_pum_id"`
	DateCreated time.Time `json:"ppk_inbox_from_pum_date_created"`
	WorkRequest WorkRequest
}

type FromPPE struct {
	ID           int64     `json:"ppk_inbox_from_ppe_id"`
	DateCreated  time.Time `json:"ppk_inbox_from_ppe_date_created"`
	EstDate      time.Time `json:"ppk_inbox_from_ppe_est_date"`
	EstLaborHour int64     `json:"ppk_inbox_from_ppe_est_labor_hour"`
	Worker       string    `json:"ppk_inbox_from_ppe_worker"`
	Worker_Email string    `json:"ppk_inbox_from_ppe_worker_email"`
	Cost         float64   `json:"ppk_inbox_from_ppe_cost"`
	WorkRequest  WorkRequest
}

type FromWorker struct {
	ID          int64     `json:"ppk_inbox_from_worker_id"`
	DateCreated time.Time `json:"ppk_inbox_from_worker_date_created"`
	WorkOrder   WorkOrder
}

func (w *FromPUM) Create() {

}

func (w *FromPPE) Create() {

}

func (w *FromWorker) Create() {

}
