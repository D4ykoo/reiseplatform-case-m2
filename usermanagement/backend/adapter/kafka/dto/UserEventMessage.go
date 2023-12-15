package dto

import domain "github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"

type UserEventMessage struct {
	UserEvent       domain.UserEvent `json:"userEvent"`
	OptionalContent string           `json:"optionalContent"`
}
