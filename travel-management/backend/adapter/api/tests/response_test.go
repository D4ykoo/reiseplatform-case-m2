package api_test

import (
	"testing"

	"github.com/mig3177/travelmanagement/adapter/api/dto"
	"github.com/mig3177/travelmanagement/domain/model"
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

func TestToResponse(t *testing.T) {
	input := hotel
	output := dto.ToHotelResoponse(input)

	if output.Id != input.Id {
		t.Fatalf(`want: %d - got: %d`, input.Id, output.Id)
	}
	if output.HotelName != input.Name {
		t.Fatalf(`want: %s - got: %s`, input.Name, output.HotelName)
	}
	if output.Description != input.Description {
		t.Fatalf(`want: %s - got: %s`, input.Description, output.Description)
	}
	if output.Land != input.Address.Land {
		t.Fatalf(`want: %s - got: %s`, input.Address.Land, output.Land)
	}
	if output.VendorName != input.Vendor.Username {
		t.Fatalf(`want: %s - got: %s`, input.Vendor.Username, output.VendorName)
	}
	if output.Tags[0].Name != input.Tags[0].Name {
		t.Fatalf(`want: %s - got: %s`, input.Tags[0].Name, output.Tags[0].Name)
	}
	if output.Pictures[0].Description != input.Pictures[0].Description {
		t.Fatalf(`want: %s - got: %s`, input.Pictures[0].Description, output.Pictures[0].Description)
	}
}
