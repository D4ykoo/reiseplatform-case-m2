package domain

// Enum functionality for PostEvents and interface description

type UserEvent int

const (
	EventUserCreate UserEvent = iota + 1
	EventUserDelete
	EventUserUpdate
	EventUserGet
	EventPasswordReset
	EventGenerateJWT
	EventLogin
	EventEmailNotification
	EventRegister
	EventLogout
)

func (userEvent UserEvent) String() string {
	return [...]string{
		"EventUserCreate",
		"EventUserDelete",
		"EventUserUpdate",
		"EventUserGet",
		"EventPasswordReset",
		"EventGenerateJWT",
		"EventLogin",
		"EventEmailNotification",
		"EventRegister",
		"EventLogout",
	}[userEvent-1]
}

func (userEvent UserEvent) EnumIndex() int {
	return int(userEvent)
}
