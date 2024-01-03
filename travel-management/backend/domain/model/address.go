package model

import "encoding/json"

type Address struct {
	Street string `json:"steet"`
	State  string `json:"state"`
	Land   string `json:"land"`
}

func (a *Address) String() string {
	res, err := json.Marshal(a)
	if err != nil {
		return ""
	}
	return string(res)
}
