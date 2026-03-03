package models

import "time"

type Task struct {
	ID				int      	`json:"seq_id_task"`
	StrTitle		string   	`json:"str_title"`
	StrDescription	string		`json:"str_description"`
	BooDone			bool     	`json:"boo_done"`
	TimCreated		time.Time	`json:"ts_created"`
	DatStart		time.Time	`json:"dat_start"`
	DatEnd			time.Time	`json:"dat_end"`
}