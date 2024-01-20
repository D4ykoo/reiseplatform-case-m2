package outbound

import (
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"
)

type IUserEvents interface {
	SendEvent(event domain.UserEvent, content string)
}
