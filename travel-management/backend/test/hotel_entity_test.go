package test

import (
	"testing"

	"github.com/mig3177/travelmanagement/adapter/dbGoRm/entities"
	"github.com/mig3177/travelmanagement/domain/model"
	"gorm.io/gorm"
)

var name = "test"
var adress = model.Address{Street: "street", State: "state", Land: "land"}
var description = "description"
var vendor = model.Vendor{Id: 55, Username: "username"}
var picture1 = model.Picture{Id: 777, Payload: "payload", Description: "des"}
var picture2 = model.Picture{Id: 888, Payload: "payload2", Description: "des2"}
var tag1 = model.Tag{Id: 477, Name: "tagname"}
var hotel = &model.Hotel{Id: 5, Name: name, Address: adress,
	Description: description, Vendor: vendor, Pictures: []*model.Picture{&picture1, &picture2}, Tags: []*model.Tag{&tag1}}

func toHotelEntity(t *testing.T) {

	output := entities.ToHotelEntity(hotel)
	if output.ID != 5 {
		t.Fatalf(`want: %d - got: %d`, 5, output.ID)
	}
	if output.Name != name {
		t.Fatalf(`want: %s - got: %s`, name, output.Name)
	}
	if output.State != adress.State {
		t.Fatalf(`want: %s - got: %s`, adress.State, output.State)
	}
	if output.Street != adress.Street {
		t.Fatalf(`want: %s - got: %s`, adress.Street, output.Street)
	}
	if output.Land != adress.Land {
		t.Fatalf(`want: %s - got: %s`, adress.Land, output.Land)
	}
	if output.Description != description {
		t.Fatalf(`want: %s - got: %s`, description, output.Description)
	}
	if output.VendorRef != vendor.Id {
		t.Fatalf(`want: %d - got: %d`, vendor.Id, output.VendorRef)
	}
	if len(output.Pictures) != 2 {
		t.Fatalf(`want: %d - got: %d`, 2, len(output.Pictures))
	}
	if len(output.Tags) != 1 {
		t.Fatalf(`want: %d - got: %d`, 1, len(output.Tags))
	}
}

func toHotelModel(t *testing.T) {
	input := &entities.HotelEntity{Model: gorm.Model{ID: 8}, Name: name, Street: adress.Street, State: adress.State, Land: adress.Land}
	output := entities.ToHotelModel(input)
	if output.Id != 8 {
		t.Fatalf(`want: %d - got: %d`, 5, output.Id)
	}
	if output.Name != name {
		t.Fatalf(`want: %s - got: %s`, name, output.Name)
	}
	if output.Address.State != adress.State {
		t.Fatalf(`want: %s - got: %s`, adress.State, output.Address.State)
	}
	if output.Address.Street != adress.Street {
		t.Fatalf(`want: %s - got: %s`, adress.Street, output.Address.Street)
	}
	if output.Description != description {
		t.Fatalf(`want: %s - got: %s`, description, output.Description)
	}
	if output.Vendor.Id != vendor.Id {
		t.Fatalf(`want: %d - got: %d`, vendor.Id, output.Vendor.Id)
	}
}
