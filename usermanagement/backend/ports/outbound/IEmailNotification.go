package outbound

type EmailContent struct {
	Header string `json:"header"`
	Title  string `json:"title"`
	From   string `json:"from"`
	To     string `json:"to"`
	Body   string `json:"body"`
}

type IEmailNotification interface {
	SendEmail(content EmailContent)
}
