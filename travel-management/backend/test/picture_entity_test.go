package test

import (
	"testing"

	"github.com/mig3177/travelmanagement/adapter/dbGoRm/entities"
	"gorm.io/gorm"
)

func toPicEntity(t *testing.T) {

	output := entities.ToPicEntity(&picture1)
	if output.ID != picture1.Id {
		t.Fatalf(`want: %d - got: %d`, picture1.Id, output.ID)
	}
	if output.Description != picture1.Description {
		t.Fatalf(`want: %s - got: %s`, picture1.Description, output.Description)
	}
	if output.Payload != picture1.Payload {
		t.Fatalf(`want: %s - got: %s`, picture1.Payload, output.Payload)
	}
}

func toPicModel(t *testing.T) {
	input := &entities.PictureEntity{Model: gorm.Model{ID: 85}, Payload: "payload", HotelRef: 55}
	output := entities.ToPicModel(input)
	if output.Id != 85 {
		t.Fatalf(`want: %d - got: %d`, 85, output.Id)
	}
	if output.Payload != picture1.Payload {
		t.Fatalf(`want: %s - got: %s`, picture1.Payload, output.Payload)
	}
	if output.Description != picture1.Description {
		t.Fatalf(`want: %s - got: %s`, output.Description, output.Description)
	}
}
