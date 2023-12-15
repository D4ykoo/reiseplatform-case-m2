package kafka

import (
	domain "github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"
)

// SendEvent Send kafka message of a user management event with some content
// Sarama package is used. Fast and efficient IBM kafka library.
//
// No return error since it does not matter, just a local panic
func SendEvent(event domain.UserEvent, content string) {
	messageString := event.String() + ": " + content
	Publish(messageString)

	return
}
