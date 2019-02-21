// Package mail provides a small wrapper around AWS SES for sending emails.
package mail

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/ses/sesiface"
)

// client is the default client.
var client = New(ses.New(session.New(aws.NewConfig())))

// Email options.
type Email struct {
	// From is the source email.
	From string

	// To is a set of destination emails.
	To []string

	// ReplyTo is a set of reply to emails.
	ReplyTo []string

	// Subject is the email subject text.
	Subject string

	// Text is the plain text representation of the body.
	Text string

	// HTML is the HTML representation of the body.
	HTML string
}

// Client is the mail client.
type Client struct {
	Service sesiface.SESAPI
}

// New returns a new mail client with the given SES service implementation.
func New(s sesiface.SESAPI) *Client {
	return &Client{
		Service: s,
	}
}

// Send an email.
func (c *Client) Send(e Email) error {
	if e.HTML == "" {
		e.HTML = e.Text
	}

	msg := &ses.Message{
		Subject: &ses.Content{
			Charset: aws.String("utf-8"),
			Data:    &e.Subject,
		},
		Body: &ses.Body{
			Html: &ses.Content{
				Charset: aws.String("utf-8"),
				Data:    &e.HTML,
			},
			Text: &ses.Content{
				Charset: aws.String("utf-8"),
				Data:    &e.Text,
			},
		},
	}

	dest := &ses.Destination{
		ToAddresses: aws.StringSlice(e.To),
	}

	_, err := c.Service.SendEmail(&ses.SendEmailInput{
		Source:           &e.From,
		Destination:      dest,
		Message:          msg,
		ReplyToAddresses: aws.StringSlice(e.ReplyTo),
	})

	return err
}

// Send an email.
func Send(e Email) error {
	return client.Send(e)
}
