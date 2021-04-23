package disgohook

import (
	"strings"

	"github.com/DisgoOrg/disgohook/api"
	"github.com/DisgoOrg/disgohook/internal"
	"github.com/DisgoOrg/log"
)

func NewWebhookByToken(logger log.Logger, webhookToken string) (api.Webhook, error) {
	webhookTokenSplit := strings.SplitN(webhookToken, "/", 2)
	if len(webhookTokenSplit) != 2 {
		return nil, api.ErrMalformedWebhookToken
	}
	return internal.NewWebhookImpl(nil, logger, webhookTokenSplit[1], webhookTokenSplit[0]), nil
}

func NewWebhookByIDToken(logger log.Logger, webhookID string, webhookToken string) (api.Webhook, error) {
	return internal.NewWebhookImpl(nil, logger, webhookID, webhookToken), nil
}
