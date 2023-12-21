package model

import "encoding/json"

type Vendor struct {
	Id       uint   `json:"id"`
	Username string `json:"name"`
}

func (v *Vendor) String() string {
	res, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(res)
}
