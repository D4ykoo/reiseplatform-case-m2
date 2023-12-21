package model

import "encoding/json"

type Hotel struct {
	Id          uint       `json:"id"`
	Name        string     `json:"name"`
	Address     Address    `json:"address"`
	Description string     `json:"description"`
	Vendor      Vendor     `json:"vendor"`
	Pictures    []*Picture `json:"pictures"`
	Travels     []*Travel  `json:"Travels"`
}

func (h *Hotel) String() string {
	res, err := json.Marshal(h)
	if err != nil {
		return ""
	}
	return string(res)
}
