package adapter

import (
	"encoding/json"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/ports"
	"log"
)

// SendEmail Mocked by sending an EventPost message via the UserEventsImpl
func SendEmail(content ports.EmailContent) {
	brokers := []string{""}
	topic := ""

	parsedContent, err := json.Marshal(content)

	if err != nil {
		log.Panic(err)
		return
	}

	SendEvent(brokers, topic, ports.EmailNotification, string(parsedContent))
}
