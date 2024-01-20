package application

import (
	"github.com/mig3177/travelmanagement/domain/model"
	"github.com/mig3177/travelmanagement/ports/outbound"
)

type TagServiceImpl struct {
	tags outbound.TagRepository
}

func NewTagService(tagrepo outbound.TagRepository) TagServiceImpl {
	return TagServiceImpl{tags: tagrepo}
}

func (service TagServiceImpl) GetTag(id uint) (*model.Tag, error) {
	return service.tags.FindByID(id)
}

func (service TagServiceImpl) ListTags() ([]*model.Tag, error) {
	return service.tags.ListAll()
}

func (service TagServiceImpl) NewTag(name string) (*model.Tag, error) {
	return service.tags.Create(&model.Tag{Name: name})
}

func (service TagServiceImpl) RemoveTag(id uint) error {
	return service.tags.Delete(id)
}

func (service TagServiceImpl) UpdateTags(tag *model.Tag) (*model.Tag, error) {
	return service.tags.Update(tag)
}
