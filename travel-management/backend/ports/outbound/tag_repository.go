package outbound

import "github.com/mig3177/travelmanagement/domain/model"

type TagRepository interface {
	Create(*model.Tag) (*model.Tag, error)
	Update(*model.Tag) (*model.Tag, error)
	Delete(uint) error
	ListAll() ([]*model.Tag, error)
	FindByID(uint) (*model.Tag, error)
	Count() (int64, error)
}
