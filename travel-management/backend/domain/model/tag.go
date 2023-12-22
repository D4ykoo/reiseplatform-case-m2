package model

import "encoding/json"

type Tag struct {
	Id   uint   `json:"typId"`
	Name string `json:"name"`
}

func (t *Tag) String() string {
	res, err := json.Marshal(t)
	if err != nil {
		return ""
	}
	return string(res)
}
