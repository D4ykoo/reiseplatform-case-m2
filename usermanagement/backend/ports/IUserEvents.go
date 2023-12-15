package ports

import domain "github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"

// TODO: rewrite interface? i guess not completely hexagonal
type IUserEvents interface {
	SendEvent(brokerUrls []string, topic string, event domain.UserEvent, content string)
}
