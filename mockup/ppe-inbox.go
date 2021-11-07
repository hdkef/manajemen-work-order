package mockup

import "manajemen-work-order/models"

func PPEInbox(id int64) models.PPEInbox {
	return models.PPEInbox{
		ID:          id,
		EstCost:     float64(id * 1000000),
		WorkRequest: WorkRequest(id),
	}
}

func PPEInboxSlice(num int, startfrom int) []models.PPEInbox {
	var PPEInboxs []models.PPEInbox
	for i := startfrom; i <= startfrom+num; i++ {
		PPEInboxs = append(PPEInboxs, PPEInbox(int64(i)))
	}
	return PPEInboxs
}
