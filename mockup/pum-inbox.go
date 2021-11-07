package mockup

import "manajemen-work-order/models"

func PUMInbox(id int64) models.PUMInbox {
	return models.PUMInbox{
		ID:          id,
		WorkRequest: WorkRequest(id),
	}
}

func PUMInboxSlice(num int, startfrom int) []models.PUMInbox {
	var PUMInboxs []models.PUMInbox
	for i := startfrom; i <= startfrom+num; i++ {
		PUMInboxs = append(PUMInboxs, PUMInbox(int64(i)))
	}
	return PUMInboxs
}
