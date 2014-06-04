package handshakejstransport

import (
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
)

var (
	SMTP_ADDRESS  string
	SMTP_PORT     string
	SMTP_USERNAME string
	SMTP_PASSWORD string
)

func Setup(smtp_address, smtp_port, smtp_username, smtp_password string) {
	SMTP_ADDRESS = smtp_address
	SMTP_PORT = smtp_port
	SMTP_USERNAME = smtp_username
	SMTP_PASSWORD = smtp_password
}

func ViaEmail(to, from, subject, text, html string) {
	e := email.NewEmail()
	e.From = from
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(text)
	e.HTML = []byte(html)

	err := e.Send(SMTP_ADDRESS+":"+SMTP_PORT, smtp.PlainAuth("", SMTP_USERNAME, SMTP_PASSWORD, SMTP_ADDRESS))
	if err != nil {
		log.Println(err)
	}
}
