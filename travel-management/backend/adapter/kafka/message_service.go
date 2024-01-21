package kafka

type MessageService interface {
	PublishAsJSON(interface{}) error
}
