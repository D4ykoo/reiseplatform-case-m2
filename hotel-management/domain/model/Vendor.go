package model

import "github.com/google/uuid"

type Vendor struct {
	ID       uuid.UUID
	Username string
}
