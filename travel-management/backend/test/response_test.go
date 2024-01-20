package test

import (
	"testing"

	"github.com/mig3177/travelmanagement/adapter/api/dto"
)

func toResponse(t *testing.T) {
	input := hotel
	output := dto.ToHotelResoponse(input)

	if output.Id != 638 {
		t.Fatalf(`want: %d - got: %d`, input.Id, 632)
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
