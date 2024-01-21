package kafka

import (
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter/kafka/dto"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"
	"log"
	"os"
	"time"
)

type UserEventImpl struct {
	MsgService MessageService
}

func Init() UserEventImpl {
	service, err := NewMessageService([]string{os.Getenv("BROKERS")})
	if err != nil {
		log.Default().Printf(err.Error())
	}

	return UserEventImpl{MsgService: service}
}

func (userEventImpl UserEventImpl) SendEvent(event domain.UserEvent, content string) {
	eventMessage := dto.UserEventMessage{
		Type: string(rune(event)),
		Log:  content,
		Time: time.Now().UTC(),
	}

	err := userEventImpl.MsgService.Publish(eventMessage)
	if err != nil {
		log.Default().Printf(err.Error())
	}
}

// SendEvent Send kafka message of a user management event with some content
// Sarama package is used. Fast and efficient IBM kafka library.
//
//// No return error since it does not matter
//func SendEventOld(event domain.UserEvent, content string) {
//	Publish(dto.UserEventMessage{
//		Type: string(rune(event)),
//		Log:  content,
//		Time: time.Now().UTC(),
//	})
//
//	return
//}
