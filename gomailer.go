// Package gomailer provides a small wrapper around AWS SES.
package gomailer

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/ses/sesiface"
)

// client is the default client.
var client = New(ses.New(session.New(aws.NewConfig())))

// Client for SES.
type Client struct {
	ses sesiface.SESAPI // Service implementation
}

// New client.
func New(ses sesiface.SESAPI) *Client {
	return &Client{
		ses: ses,
	}
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
func (c *Client) SendEmail(email Email) error {
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

	_, err := c.ses.SendEmail(&ses.SendEmailInput{
		Source:           &email.From,
		Destination:      dest,
		Message:          msg,
		ReplyToAddresses: aws.StringSlice(email.ReplyTo),
	})

	return err
}

// SendEmail message.
func SendEmail(email Email) error {
	return client.SendEmail(email)
}
