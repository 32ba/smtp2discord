package main

import (
	"io"
	"net/mail"
	"time"
)

type Email struct {
    To                  *[]*mail.Address
    From                *[]*mail.Address
    Subject             *string
    Body                *string
    Date                *time.Time
    ContentType         *string
    ContentTransferType *string
}

func parseEmail(msg *mail.Message) (*Email, error) {
    to, err := msg.Header.AddressList("To")
    if err != nil {
        return nil, err
    }
	from, err := msg.Header.AddressList("From")
	if err != nil {
		return nil, err
	}
	date, err := mail.ParseDate(msg.Header.Get("Date"))
	if err != nil {
		return nil, err
	}
	subject := msg.Header.Get("Subject")
	contentType := msg.Header.Get("Content-Type")
	contentTransferType := msg.Header.Get("Content-Transfer-Encoding")
	
	bodyBytes, err := io.ReadAll(msg.Body)
	if err != nil {
		return nil, err
	}
	body := string(bodyBytes)
	
	email := &Email{
		To:                  &to,
		From:                &from,
		Subject:             &subject,
		Body:                &body,
		Date:                &date,
		ContentType:         &contentType,
		ContentTransferType: &contentTransferType,
	}
	
	return email, nil
}
