package mockup

import (
	"fmt"
	"manajemen-work-order/models"
	"time"
)

func WorkOrder(id int64) models.WorkOrder {
	return models.WorkOrder{
		ID:           id,
		WorkRequest:  WorkRequest(id),
		EstDate:      time.Now(),
		EstLaborHour: id,
		Worker:       fmt.Sprintf("Worker %d", id),
		Worker_Email: fmt.Sprintf("worker%d@email.com", id),
		Cost:         float64(id * 1000000),
		Status:       1,
		ApproverID:   1,
	}
}

func WorkOrderSlice(num int, startfrom int) []models.WorkOrder {
	var workorders []models.WorkOrder
	for i := startfrom; i <= startfrom+num; i++ {
		workorders = append(workorders, WorkOrder(int64(i)))
	}
	return workorders
}
