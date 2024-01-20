package domain

import (
	"github.com/mig3177/travelmanagement/domain/model"
)

type TagService interface {
	NewTag(string) (*model.Tag, error)
	GetTag(uint) (*model.Tag, error)
	ListTags() ([]*model.Tag, error)
	RemoveTag(uint) error
	UpdateTags(*model.Tag) (*model.Tag, error)
}
