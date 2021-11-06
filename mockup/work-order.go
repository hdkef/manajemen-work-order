package mockup

import (
	"manajemen-work-order/models"
	"time"
)

var WORK_ORDER models.WorkOrder = models.WorkOrder{
	ID:           1,
	WorkRequest:  WORK_REQUEST,
	EstDate:      time.Date(2021, 12, 2, 8, 0, 0, 0, time.Now().Location()),
	EstLaborHour: 5,
	Worker:       "PT ABADI",
	Worker_Email: "hdkef11@gmail.com",
	Cost:         300000,
	Status:       "ON PROGRESS",
	ApproverID:   2,
}
