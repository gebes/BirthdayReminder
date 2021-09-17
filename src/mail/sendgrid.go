package mail

import (
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"gebes.io/go-birthdayreminder/src/env"
)



func SendGridMail(mailBody* Mail) (*rest.Response, error) {
	from := mail.NewEmail("Birthday Reminder", "noreply@gebes.io")
	subject := mailBody.Subject
	to := mail.NewEmail(mailBody.Receiver, mailBody.Receiver)
	message := mail.NewSingleEmail(from, subject, to, "", mailBody.Body)
	client := sendgrid.NewSendClient(env.SendgridApiKey)
	return client.Send(message)
}
