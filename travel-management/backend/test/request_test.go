package test

import (
	"testing"

	"github.com/mig3177/travelmanagement/adapter/api/dto"
)

func toRequest(t *testing.T) {
	input := dto.UpdateHotelRequest{Id: 632, HotelName: "hotel", Street: "street", State: "state", Land: "land",
		VendorId: 88, VendorName: "vendorname", Description: "description",
		Pictures: []dto.PictureRequest{{Id: 784, Description: "picDescription", Payload: "payload"}},
		Tags:     []dto.TagRequest{{Id: 5852, Name: "tag"}}}

	output := dto.ToHotelModel(&input)
	if output.Id != 632 {
		t.Fatalf(`want: %d - got: %d`, input.Id, 632)
	}
	if output.Name != input.HotelName {
		t.Fatalf(`want: %s - got: %s`, input.HotelName, output.Name)
	}
	if output.Description != input.Description {
		t.Fatalf(`want: %s - got: %s`, input.Description, output.Description)
	}
	if output.Address.Land != input.Land {
		t.Fatalf(`want: %s - got: %s`, input.Land, output.Address.Land)
	}
	if output.Vendor.Username != input.VendorName {
		t.Fatalf(`want: %s - got: %s`, input.VendorName, output.Vendor.Username)
	}
	if output.Tags[0].Name != input.Tags[0].Name {
		t.Fatalf(`want: %s - got: %s`, input.Tags[0].Name, output.Tags[0].Name)
	}
	if output.Pictures[0].Description != input.Pictures[0].Description {
		t.Fatalf(`want: %s - got: %s`, input.Pictures[0].Description, output.Pictures[0].Description)
	}
}
