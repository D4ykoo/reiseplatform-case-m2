package dto

import "time"

type UserEventMessage struct {
	Type string    `json:"type"`
	Log  string    `json:"log"`
	Time time.Time `json:"time"`
}
