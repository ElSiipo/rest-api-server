package models

type Instruction struct {
	ID          string `db:"id" json:"id"`
	EventStatus string `db:"event_status" json:"event_status"`
	EventName   string `db:"event_name" json:"event_name"`
}
