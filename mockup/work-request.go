package mockup

import (
	"manajemen-work-order/models"
	"time"
)

var WORK_REQUEST models.WorkRequest = models.WorkRequest{
	ID:          1,
	RequestorID: 1,
	Priority:    1,
	DateCreated: time.Now(),
	Task:        "Memperbaiki pipa 2\"",
	Equipment:   "Boiler 2",
	Location:    "Utilities",
	Instruction: "tes kebocoran dan mengganti pipa kemudian cek kembali",
	Description: "",
	Status:      1,
}
