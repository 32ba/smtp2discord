package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alash3al/go-smtpsrv/v3"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "2525"
	}
	if os.Getenv("WEBHOOK_URL") == "" {
		log.Fatal("Error: WEBHOOK_URL is not set")
	}

	cfg := smtpsrv.ServerConfig{
		BannerDomain:    "localhost",
		ListenAddr:      fmt.Sprintf(":%s", port),
		MaxMessageBytes: 1024 * 1024 * 10,
		ReadTimeout:     time.Duration(10) * time.Second,
		WriteTimeout:    time.Duration(10) * time.Second,
		Handler:         Handler,
	}

	log.Printf("Info: Starting SMTP server on port %s", port)

	err := smtpsrv.ListenAndServe(&cfg)
	if err != nil {
		log.Fatalf("Error: Failed to start SMTP server: %v", err)
	}
}
