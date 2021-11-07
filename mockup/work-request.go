package mockup

import (
	"fmt"
	"manajemen-work-order/models"
	"time"
)

func WorkRequest(id int64) models.WorkRequest {
	return models.WorkRequest{
		ID:          id,
		RequestorID: 1,
		Priority:    1,
		DateCreated: time.Now(),
		Task:        fmt.Sprintf("Task %d", id),
		Equipment:   fmt.Sprintf("equipment %d", id),
		Location:    fmt.Sprintf("Location %d", id),
		Instruction: fmt.Sprintf("Instruction %d", id),
		Description: fmt.Sprintf("Description %d", id),
		Status:      1,
	}
}

func WorkRequestSlice(num int, startfrom int) []models.WorkRequest {
	var workrequests []models.WorkRequest
	for i := startfrom; i <= startfrom+num; i++ {
		workrequests = append(workrequests, WorkRequest(int64(i)))
	}
	return workrequests
}
