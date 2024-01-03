package model

import (
	"encoding/json"
	"time"
)

type Travel struct {
	Id          uint      `json:"id"`
	Vendor      Vendor    `json:"vendor"`
	From        time.Time `json:"from"`
	To          time.Time `json:"to"`
	Price       float32   `json:"price"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (t *Travel) String() string {
	res, err := json.Marshal(t)
	if err != nil {
		return ""
	}
	return string(res)
}
