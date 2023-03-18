package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/alash3al/go-smtpsrv/v3"
)

func Handler(c *smtpsrv.Context) error {
	log.Printf("Info: Received message from %s", c.RemoteAddr().String())
	webhookUrl := os.Getenv("WEBHOOK_URL")

	msg, err := c.Parse()
	if err != nil {
		return handleError(fmt.Errorf("failed to parse message: %v", err))
	}

	webhookMessage := messageBuilder(msg)

	json, err := json.Marshal(webhookMessage)
	if err != nil {
		return handleError(fmt.Errorf("failed to marshal message: %v", err))
	}

	resp, err := send(json, webhookUrl)
	if err != nil || resp.StatusCode != 200 {
		return handleError(fmt.Errorf("failed to send message: %v", err))
	}

	log.Printf("Info: Message sent successfully")

	return nil
}

func messageBuilder(msg *smtpsrv.Email) *DiscordWebhook {
	var toStr string = ""
	for _, s := range msg.To {
		if s.Name != "" {
			toStr += fmt.Sprintf(" %s <%s>", s.Name, s.Address)
		} else {
			toStr += fmt.Sprintf(" %s", s.Address)
		}
	}

	var fromStr string = ""
	for _, s := range msg.From {
		if s.Name != "" {
			fromStr += fmt.Sprintf(" %s <%s>", s.Name, s.Address)
		} else {
			fromStr += fmt.Sprintf(" %s", s.Address)
		}
	}

	return &DiscordWebhook{
		Content: fmt.Sprintf("New email received from: %s", fromStr),
		Embeds: DiscordEmbed{
			{
				Title:       fmt.Sprintf("Subject: %s", msg.Subject),
				Description: fmt.Sprintf("```%s```", msg.TextBody),
				Color:       0x00ff00,
				Author: DiscordEmbedAuthor{
					Name:   "Email",
				},
				Fields: []DiscordEmbedField{
					{
						Name:   "From",
						Value:  fromStr,
						Inline: true,
					},
					{
						Name:   "To",
						Value:  toStr,
						Inline: true,
					},
					{
						Name:   "Date",
						Value:  msg.Date.String(),
						Inline: true,
					},
				},
			},
		},
	}
}

func send(json []byte, url string) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(json))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("wait", "true")
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}

	return resp, nil
}

func handleError(err error) error {
	log.Printf("Error: %v", err)
	return err
}
