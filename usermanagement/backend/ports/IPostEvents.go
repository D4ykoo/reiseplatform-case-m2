package ports

// Enum functionality for PostEvents and interface description

type PostEvent int

const (
	UserCreate PostEvent = iota
	UserDelete
	UserUpdate
	UserGet
	PasswordReset
	GenerateJWT
	Login
	EmailNotification
)

func (postEvent PostEvent) String() string {
	return [...]string{
		"UserCreate",
		"UserDelete",
		"UserUpdate",
		"UserGet",
		"PasswordReset",
		"GenerateJWT",
		"Login",
		"EmailNotification"}[postEvent+1]
}

func (postEvent PostEvent) EnumIndex() int {
	return int(postEvent)
}

// TODO: rewrite interface? i guess not completely hexagonal
type IPostEvents interface {
	initProducer(brokerUrls []string)
	SendEvent(brokerUrls []string, topic string, event PostEvent, content string)
}
