package kafka

import (
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter/kafka/dto"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"
)

// SendEvent Send kafka message of a user management event with some content
// Sarama package is used. Fast and efficient IBM kafka library.
//
// No return error since it does not matter
func SendEvent(event domain.UserEvent, content string) {
	Publish(dto.UserEventMessage{
		UserEvent:       event,
		OptionalContent: content,
	})

	return
}
