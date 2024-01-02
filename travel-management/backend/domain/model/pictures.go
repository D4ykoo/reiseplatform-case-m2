package model

import "encoding/json"

type Picture struct {
	Id          uint   `json:"id"`
	Payload     string `json:"payload"`
	Description string `json:"description"`
}

func (p *Picture) String() string {
	res, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(res)
}
