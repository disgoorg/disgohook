package disgohook

import (
	"strings"

	"github.com/DisgoOrg/disgohook/api"
	"github.com/DisgoOrg/disgohook/internal"
	"github.com/DisgoOrg/log"
)

func NewWebhookByToken(webhookToken string, logger log.Logger) (api.Webhook, error) {
	webhookTokenSplit := strings.SplitN(webhookToken, "/", 2)
	if len(webhookTokenSplit) != 2 {
		return nil, api.ErrMalformedWebhookToken
	}
	return internal.NewWebhookImpl(nil, logger, webhookTokenSplit[0], webhookTokenSplit[1]), nil
}

func NewWebhookByID(id string, token string) (api.Webhook, error) {
	return nil, nil
}
