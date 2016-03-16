// Package ses provides a small wrapper around AWS SES.
package ses

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/ses/sesiface"
)

// SES alert support.
type SES struct {
	Service sesiface.SESAPI // Service implementation
}

// Email options.
type Email struct {
	From    string   // From source email
	To      []string // To destination email(s)
	Subject string   // Subject text to send
	Text    string   // Text is the text body representation
	HTML    string   // HTMLBody is the HTML body representation
	ReplyTo []string // Reply-To email(s)
}

// SendEmail message.
func (s *SES) SendEmail(email Email) error {
	if email.HTML == "" {
		email.HTML = email.Text
	}

	msg := &ses.Message{
		Subject: &ses.Content{
			Charset: aws.String("utf-8"),
			Data:    &email.Subject,
		},
		Body: &ses.Body{
			Html: &ses.Content{
				Charset: aws.String("utf-8"),
				Data:    &email.HTML,
			},
			Text: &ses.Content{
				Charset: aws.String("utf-8"),
				Data:    &email.Text,
			},
		},
	}

	dest := &ses.Destination{
		ToAddresses: aws.StringSlice(email.To),
	}

	_, err := s.Service.SendEmail(&ses.SendEmailInput{
		Source:           &email.From,
		Destination:      dest,
		Message:          msg,
		ReplyToAddresses: aws.StringSlice(email.ReplyTo),
	})

	return err
}
