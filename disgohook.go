package disgohook

import (
	"strings"

	"github.com/DisgoOrg/disgohook/api"
	"github.com/DisgoOrg/disgohook/internal"
)

func NewWebhookByToken(webhookToken string) (api.Webhook, error) {
	webhookTokenSplit := strings.SplitN(webhookToken, "/", 2)
	if len(webhookTokenSplit) != 2 {
		return nil, api.ErrMalformedWebhookToken
	}
	return internal.NewWebhookImpl(webhookTokenSplit[0], webhookTokenSplit[1]), nil
}

func NewWebhookByID(id string, token string) (api.Webhook, error) {
	return nil, nil
}
