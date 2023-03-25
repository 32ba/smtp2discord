package main

import (
	"encoding/base64"
	"io"
	"mime"
	"mime/quotedprintable"
	"net/mail"
	"strings"
	"time"
)

type Email struct {
	To                      []*mail.Address
	From                    []*mail.Address
	Subject                 string
	Body                    string
	Date                    time.Time
	MIMEVersion             string
	ContentType             string
	ContentTransferEncoding string
}

func parseEmail(msg *mail.Message) (*Email, error) {
	to, err := msg.Header.AddressList("To")
	if err != nil {
		to = []*mail.Address{
			{
				Name: "Anonymous",
			},
		}
	}
	from, err := msg.Header.AddressList("From")
	if err != nil {
		from = []*mail.Address{
			{
				Name: "Anonymous",
			},
		}
	}
	date, err := mail.ParseDate(msg.Header.Get("Date"))
	if err != nil {
		date = time.Unix(0, 0)
	}
	subject := msg.Header.Get("Subject")
	decodedSubject, _ := decodeSubject(subject)
	mimeVersion := msg.Header.Get("MIME-Version")
	contentType := msg.Header.Get("Content-Type")
	contentTransferEncoding := msg.Header.Get("Content-Transfer-Encoding")

	bodyBytes, err := io.ReadAll(msg.Body)
	if err != nil {
		bodyBytes = []byte{}
	}
	body := string(bodyBytes)
	decodedBody, _ := decodeBody(body, contentTransferEncoding)

	email := &Email{
		To:                      to,
		From:                    from,
		Subject:                 decodedSubject,
		Body:                    decodedBody,
		Date:                    date,
		MIMEVersion:             mimeVersion,
		ContentType:             contentType,
		ContentTransferEncoding: contentTransferEncoding,
	}

	return email, nil
}

func decodeSubject(encodedSubject string) (string, error) {
	if !strings.Contains(encodedSubject, "=?") && !strings.Contains(encodedSubject, "?=") {
		return encodedSubject, nil
	}

	dec := new(mime.WordDecoder)
	dec.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch strings.ToLower(charset) {
		case "utf-8", "iso-8859-1", "iso8859-1", "us-ascii":
			return input, nil
		default:
			return nil, nil
		}
	}
	return dec.DecodeHeader(encodedSubject)
}

func decodeBody(encodedBody string, contentTransferEncoding string) (string, error) {
	switch strings.ToLower(contentTransferEncoding) {
	case "base64":
		decodedBody, err := base64.StdEncoding.DecodeString(encodedBody)
		if err != nil {
			return "", err
		}
		return string(decodedBody), nil
	case "quoted-printable":
		decoder := quotedprintable.NewReader(strings.NewReader(encodedBody))
		decodedBody, err := io.ReadAll(decoder)
		if err != nil {
			return "", err
		}
		return string(decodedBody), nil
	default:
		return encodedBody, nil
	}
}
