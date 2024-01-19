package test

import (
	"testing"

	"github.com/mig3177/travelmanagement/adapter/dbGoRm/entities"
	"gorm.io/gorm"
)

func toTagEntity(t *testing.T) {

	output := entities.ToTagEntity(&tag1)
	if output.ID != tag1.Id {
		t.Fatalf(`want: %d - got: %d`, tag1.Id, output.ID)
	}
	if output.Name != tag1.Name {
		t.Fatalf(`want: %s - got: %s`, tag1.Name, output.Name)
	}
}

func toTagModel(t *testing.T) {
	input := &entities.TagEntity{Model: gorm.Model{ID: 88}, Name: tag1.Name}
	output := entities.ToTagModel(input)
	if output.Id != 88 {
		t.Fatalf(`want: %d - got: %d`, 88, output.Id)
	}
	if output.Name != tag1.Name {
		t.Fatalf(`want: %s - got: %s`, tag1.Name, output.Name)
	}
}
