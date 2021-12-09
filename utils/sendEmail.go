package utils

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"
)

func SendEmail(text string, to_email string) {
	// Sender data.
	from := "vardbax1995@mail.ru"
	password := "4A9bX3ZQc39asx2puXAq"

	// Receiver email address.
	to := []string{
		to_email,
	}

	// smtp server configuration.
	smtpHost := "smtp.mail.ru"
	smtpPort := "587"

	// Message.
	message := []byte(text)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}

type Email struct {
	Price    string
	Location string
	Link     string
}

func ParseEmailTemplate(content []string) string {

	t, _ := template.ParseFiles("./template/email.html")
	buf := new(bytes.Buffer)
	t.Execute(buf, content)
	return buf.String()

}
