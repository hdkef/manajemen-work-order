package models

type SPKProgress struct {
	Status string `json:"status"`
	PIN    int64  `json:"pin"`
}
