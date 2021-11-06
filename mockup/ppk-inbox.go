package mockup

import (
	"manajemen-work-order/models"
	"time"
)

var PPK_INBOX_FROM_PUM models.PPKInbox = models.PPKInbox{
	FromPUM: models.FromPUM{
		ID:          1,
		DateCreated: time.Now(),
		WorkRequest: WORK_REQUEST,
	},
}

var PPK_INBOX_FROM_PPE models.PPKInbox = models.PPKInbox{
	FromPPE: models.FromPPE{
		ID:           1,
		DateCreated:  time.Now(),
		EstDate:      time.Date(2021, 12, 2, 8, 0, 0, 0, nil),
		EstLaborHour: 5,
		Worker:       "PT ABADI",
		Worker_Email: "hdkef11@gmail.com",
		Cost:         300000,
		WorkRequest:  WORK_REQUEST,
	},
}

var PPK_INBOX_FROM_WORKER models.PPKInbox = models.PPKInbox{
	FromWorker: models.FromWorker{
		ID:          1,
		DateCreated: time.Now(),
		WorkOrder:   WORK_ORDER,
	},
}

var PPK_INBOX_FROM_PUM_SLICE []models.FromPUM = []models.FromPUM{
	PPK_INBOX_FROM_PUM.FromPUM, PPK_INBOX_FROM_PUM.FromPUM, PPK_INBOX_FROM_PUM.FromPUM,
}

var PPK_INBOX_FROM_PPE_SLICE []models.FromPPE = []models.FromPPE{
	PPK_INBOX_FROM_PPE.FromPPE, PPK_INBOX_FROM_PPE.FromPPE, PPK_INBOX_FROM_PPE.FromPPE,
}

var PPK_INBOX_FROM_WORKER_SLICE []models.FromWorker = []models.FromWorker{
	PPK_INBOX_FROM_WORKER.FromWorker, PPK_INBOX_FROM_WORKER.FromWorker, PPK_INBOX_FROM_WORKER.FromWorker,
}
