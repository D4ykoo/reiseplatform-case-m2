package model

import "github.com/google/uuid"

type Picture struct {
	ID          uuid.UUID
	Payload     string
	Description string
}
