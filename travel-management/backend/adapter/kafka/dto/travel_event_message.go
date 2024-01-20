package dto

import "time"

type HotelTravelEvent struct {
	Type string    `json:"type"`
	Log  string    `json:"log"`
	Time time.Time `json:"time"`
}
