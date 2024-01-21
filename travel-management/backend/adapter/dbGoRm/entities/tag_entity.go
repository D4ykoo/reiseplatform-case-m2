package entities

import (
	"encoding/json"

	"github.com/mig3177/travelmanagement/domain/model"
	"gorm.io/gorm"
)

type TagEntity struct {
	gorm.Model
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
	return &TagEntity{Model: gorm.Model{ID: tag.Id}, Name: tag.Name}
}

func ToTagModel(tag *TagEntity) *model.Tag {
	return &model.Tag{Id: tag.ID, Name: tag.Name}
}
