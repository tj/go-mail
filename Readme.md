# go-ses

Package ses provides a small wrapper around AWS SES.

## Badges

[![GoDoc](https://godoc.org/github.com/tj/go-ses?status.svg)](https://godoc.org/github.com/tj/go-ses)
![](https://img.shields.io/badge/license-MIT-blue.svg)
![](https://img.shields.io/badge/status-stable-green.svg)

## Basic Usage

```` go
package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	sses "github.com/tj/go-ses"
)

func main() {
	svc := ses.New(session.New(aws.NewConfig()))
	mailer := sses.SES{svc}

	email := sses.Email{
		From:    "me@example.com",
		To:      []string{"you@example.com"},
		Subject: "Hey you!",
		Text:    "Hi",
		HTML:    "<strong>Hi</strong>",
	}

	mailer.SendEmail(email)
}
````

---

> [tjholowaychuk.com](http://tjholowaychuk.com) &nbsp;&middot;&nbsp;
> GitHub [@tj](https://github.com/tj) &nbsp;&middot;&nbsp;
> Twitter [@tjholowaychuk](https://twitter.com/tjholowaychuk)
