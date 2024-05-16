package main

import (
	"log"
	"net/smtp"
	"os"
)

// func main() {
// 	m := mail.NewMsg()
// 	if err := m.From("toni.sender@example.com"); err != nil {
// 		log.Fatalf("failed to set From address: %s", err)
// 	}
// 	if err := m.To("tina.recipient@example.com"); err != nil {
// 		log.Fatalf("failed to set To address: %s", err)
// 	}
// 	m.Subject("This is my first mail with go-mail!")
// 	m.SetBodyString(mail.TypeTextPlain, "Do you like this mail? I certainly do!")
// 	c, err := mail.NewClient("sandbox.smtp.mailtrap.io", mail.WithPort(2525), mail.WithSMTPAuth(mail.SMTPAuthPlain),
// 		mail.WithUsername(os.Getenv("SMTP_USERNAME")), mail.WithPassword(os.Getenv("SMTP_PASSWORD")))
// 	if err != nil {
// 		log.Fatalf("failed to create mail client: %s", err)
// 	}
// 	if err := c.DialAndSend(m); err != nil {
// 		log.Fatalf("failed to send mail: %s", err)
// 	}
// 	log.Printf("Test mail successfully delivered.")
// }

var (
	from       = "gopher@example.net"
	msg        = []byte("dummy message")
	recipients = []string{"foo@example.com"}
)

func main() {
	hostname := "sandbox.smtp.mailtrap.io:25"
	auth := smtp.PlainAuth("", os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"), hostname)

	err := smtp.SendMail(hostname, auth, from, recipients, msg)
	if err != nil {
		log.Fatal(err)
	}
}
