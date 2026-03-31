package models

import "time"

type Task struct {
	ID             int       `json:"SeqIdTask"`
	StrTitle       string    `json:"StrTitle"`
	StrDescription string    `json:"StrDescription"`
	BooDone        bool      `json:"BooDone"`
	TimCreated     time.Time `json:"TimCreated"`
	DatStart       time.Time `json:"DatStart"`
	DatEnd         time.Time `json:"DatEnd"`
}
