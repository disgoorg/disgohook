package main

import (
	"github.com/DisgoOrg/disgohook/api"
	"os"

	"github.com/DisgoOrg/disgohook"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.Info("starting example...")

	webhook, err := disgohook.NewWebhookClientByToken(nil, logger, os.Getenv("webhook_token"))
	if err != nil {
		logger.Errorf("failed to create webhook: %s", err)
		return
	}

	reader, _ := os.Open("gopher.png")
	if _, err = webhook.SendMessage(api.NewWebhookMessageCreateBuilder().
		SetContent("example message").
		AddFile("gopher.png", reader).
		Build(),
	); err != nil {
		logger.Errorf("failed to send webhook message: %s", err)
		return
	}
}
