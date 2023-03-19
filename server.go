package main

import (
	"io"
	"log"
	"net/mail"

	"github.com/emersion/go-smtp"
)

type backend struct {}

func (bkd *backend) Login(state *smtp.ConnectionState, username, password string) (smtp.Session, error) {
    return &session{ state }, nil
}

func (bkd *backend) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
    return &session{ state }, nil
}

type session struct {
    ConnectionState *smtp.ConnectionState
}

func (s *session) Mail(from string, opts smtp.MailOptions) error {
    return nil
}

func (s *session) Rcpt(to string) error {
    return nil
}

func (s *session) Data(r io.Reader) error {
    log.Printf("Info: Received email from %s", s.ConnectionState.RemoteAddr)
    msg, err := mail.ReadMessage(r)
    if err != nil {
        return err
    }

	err = handler(msg)
	if err != nil {
		return err
	}

    return nil
}

func (s *session) Logout() error {
	return nil
}

func (s *session) Reset() {}