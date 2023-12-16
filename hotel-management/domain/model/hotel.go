package model

import "github.com/google/uuid"

type Hotel struct {
	ID          uuid.UUID
	Name        string
	Address     Address
	Description string
	Vendor      Vendor
	Pictures    []*Picture
}
