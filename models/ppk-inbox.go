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
	WorkOrder   WorkOrder
}

type FromPPE struct {
	ID          int64     `json:"ppk_inbox_from_ppe_id"`
	DateCreated time.Time `json:"ppk_inbox_from_ppe_date_created"`
	WorkOrder   WorkOrder
}

type FromWorker struct {
	ID          int64     `json:"ppk_inbox_from_worker_id"`
	DateCreated time.Time `json:"ppk_inbox_from_worker_date_created"`
	WorkOrder   WorkOrder
}
