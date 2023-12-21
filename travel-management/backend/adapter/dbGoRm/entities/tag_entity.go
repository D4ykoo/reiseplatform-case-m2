package entities

import (
	"encoding/json"

	"github.com/mig3177/travelmanagement/domain/model"
	"gorm.io/gorm"
)

type TagEntity struct {
	gorm.Model
	Typ     int            `json:"typ" gorm:"primary_key;"`
	Name    string         `json:"name"`
	Travels []TravelEntity `json:"travels" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;many2many:travel_tags;"`
}

func (t *TagEntity) String() string {
	res, err := json.Marshal(t)
	if err != nil {
		return ""
	}
	return string(res)
}

func ToTagEntity(tag *model.Tag) *TagEntity {
	return &TagEntity{Typ: tag.Typ, Name: tag.Name}
}

func ToTagModel(tag *TagEntity) *model.Tag {
	return &model.Tag{Typ: tag.Typ, Name: tag.Name}
}
