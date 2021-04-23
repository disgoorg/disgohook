package main

import (
	"os"

	"github.com/DisgoOrg/disgohook"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.Info("starting test...")
	webhook, err := disgohook.NewWebhookByToken(os.Getenv("webhook_token"), logger)
	if err != nil {
		logger.Errorf("failed to create webhook: %s", err)
		return
	}
	_, err = webhook.SendContent("test")
	if err != nil {
		logger.Errorf("failed to send webhook message: %s", err)
		return
	}
}