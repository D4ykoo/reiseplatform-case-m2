package adapter

import (
	"encoding/json"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter/kafka"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/ports/outbound"
	"log"
)

// SendEmail Mocked by sending an EventPost message via the UserEventsImpl
func SendEmail(content outbound.EmailContent) {
	parsedContent, err := json.Marshal(content)

	if err != nil {
		log.Panic(err)
		return
	}

	kafka.SendEvent(domain.EventEmailNotification, string(parsedContent))
}
