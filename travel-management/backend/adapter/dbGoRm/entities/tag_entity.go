package entities

import "fmt"

type TagEntity struct {
	Typ     int            `json:"typ" gorm:"primary_key;"`
	Name    string         `json:"name" `
	Travels []TravelEntity `gorm:"many2many:travel_tags;"`
}

func (b *TagEntity) String() string {
	return fmt.Sprintf("%s", b.Name)
}
