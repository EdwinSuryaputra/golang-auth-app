package smtp

type SendEmailPayload struct {
	Host       string
	Port       string
	User       string
	Password   string
	Recipients []string
	Subject    string
	HTMLBody   string
}
