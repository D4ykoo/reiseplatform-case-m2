package ports

type EmailContent struct {
	header string
	title  string
	from   string
	to     string
	body   string
}

type IEmailNotification interface {
	sendEmail(content EmailContent)
}
