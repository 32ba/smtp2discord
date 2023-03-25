package main

import (
	"log"
	"os"
	"time"

	"github.com/emersion/go-smtp"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "2525"
	}
	if os.Getenv("WEBHOOK_URL") == "" {
		log.Fatal("Error: WEBHOOK_URL is not set")
	}

	be := &backend{}
	s := smtp.NewServer(be)
	s.Addr = ":" + port
	s.Domain = "localhost"
	s.ReadTimeout = 10 * time.Minute
	s.WriteTimeout = 10 * time.Minute
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 50
	s.AllowInsecureAuth = true

	log.Printf("Info: Starting server on port %s", port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
