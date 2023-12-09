package model

import (
	"time"

	"github.com/google/uuid"
)

type Travel struct {
	ID          uuid.UUID
	Hotel       Hotel
	Vendor      Vendor
	From        time.Time
	To          time.Time
	Price       float32
	Description string
	Created     time.Time
	Edited      time.Time
	Tags        []Tag
}
