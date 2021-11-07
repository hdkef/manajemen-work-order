package mockup

import (
	"fmt"
	"manajemen-work-order/models"
	"time"
)

func PPKInboxFromPUM(id int64) models.FromPUM {
	return models.FromPUM{
		ID:          id,
		DateCreated: time.Now(),
		WorkRequest: WorkRequest(id),
	}
}

func PPKInboxFromPPE(id int64) models.FromPPE {
	return models.FromPPE{
		ID:           id,
		DateCreated:  time.Now(),
		EstDate:      time.Now(),
		EstLaborHour: id,
		Worker:       fmt.Sprintf("worker %d", id),
		Worker_Email: fmt.Sprintf("worker%d@email.com", id),
		Cost:         float64(id * 1000000),
		WorkRequest:  WorkRequest(id),
	}
}

func PPKInboxFromWorker(id int64) models.FromWorker {
	return models.FromWorker{
		ID:          id,
		DateCreated: time.Now(),
		WorkOrder:   WorkOrder(id),
	}
}

func PPKInboxFromPPESlice(num int, startfrom int) []models.FromPPE {
	var FromPPEs []models.FromPPE
	for i := startfrom; i <= startfrom+num; i++ {
		FromPPEs = append(FromPPEs, PPKInboxFromPPE(int64(i)))
	}
	return FromPPEs
}

func PPKInboxFromWorkerSlice(num int, startfrom int) []models.FromWorker {
	var FromWorkers []models.FromWorker
	for i := startfrom; i <= startfrom+num; i++ {
		FromWorkers = append(FromWorkers, PPKInboxFromWorker(int64(i)))
	}
	return FromWorkers
}

func PPKInboxFromPUMSlice(num int, startfrom int) []models.FromPUM {
	var FromPUMs []models.FromPUM
	for i := startfrom; i <= startfrom+num; i++ {
		FromPUMs = append(FromPUMs, PPKInboxFromPUM(int64(i)))
	}
	return FromPUMs
}
