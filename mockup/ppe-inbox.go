package mockup

import "manajemen-work-order/models"

var PPE_INBOX models.PPEInbox = models.PPEInbox{
	ID:          1,
	EstCost:     300000,
	WorkRequest: WORK_REQUEST,
}

var PPE_INBOX_SLICE []models.PPEInbox = []models.PPEInbox{
	PPE_INBOX, PPE_INBOX, PPE_INBOX,
}
