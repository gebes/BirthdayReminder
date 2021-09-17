package mail

import (
	"crypto/tls"
	"fmt"
	"gebes.io/go-birthdayreminder/src/env"
	gomail "gopkg.in/mail.v2"
)


// Deprecated: most of these E-Mails go into spam. Use [SendGridMail] instead
func SendMail(mail *Mail) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", env.MailServerUsername)

	// Set E-Mail receivers
	m.SetHeader("To", mail.Receiver)

	// Set E-Mail subject
	m.SetHeader("Subject", mail.Subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", mail.Body)

	// Settings for SMTP server
	d := gomail.NewDialer(env.MailServerAddress, env.MailServerPort, env.MailServerUsername, env.MailServerPassword)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: false, ServerName: env.MailServerAddress}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
