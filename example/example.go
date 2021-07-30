package main

import (
	"fmt"
	"github.com/DisgoOrg/disgohook"
	"github.com/DisgoOrg/disgohook/api"
	"os"
)

func main() {
	webhook, err := disgohook.NewWebhookClientByToken(nil, nil, os.Getenv("webhook_token"))
	if err != nil {
		fmt.Printf("failed to create webhook: %s", err)
		return
	}

	reader, _ := os.Open("gopher.png")
	if _, err = webhook.SendMessage(api.NewWebhookMessageCreateBuilder().
		SetContent("example message").
		AddFile("gopher.png", reader).
		Build(),
	); err != nil {
		fmt.Printf("failed to send webhook message: %s", err)
		return
	}
}
