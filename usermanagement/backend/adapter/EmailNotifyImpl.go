package adapter

import (
	"encoding/json"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/ports/outbound"
	"log"
)

type EmailNotifyImpl struct {
	userEvents outbound.IUserEvents
}

func InitEmail(userEvents outbound.IUserEvents) EmailNotifyImpl {
	return EmailNotifyImpl{userEvents: userEvents}
}

// SendEmail Mocked by sending an EventPost message via the UserEventsImpl
func (email EmailNotifyImpl) SendEmail(content outbound.EmailContent) {
	parsedContent, err := json.Marshal(content)

	if err != nil {
		log.Panic(err)
		return
	}

	email.userEvents.SendEvent(domain.EventEmailNotification, string(parsedContent))
}
