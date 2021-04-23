package disgohook

import (
	"net/http"
	"strings"

	"github.com/DisgoOrg/disgohook/api"
	"github.com/DisgoOrg/disgohook/internal"
	"github.com/DisgoOrg/log"
)

const Version = ""

func NewWebhookByToken(client *http.Client, logger log.Logger, webhookToken string) (api.Webhook, error) {
	webhookTokenSplit := strings.SplitN(webhookToken, "/", 2)
	if len(webhookTokenSplit) != 2 {
		return nil, api.ErrMalformedWebhookToken
	}
	return internal.NewWebhookImpl(client, logger, webhookTokenSplit[0], webhookTokenSplit[1]), nil
}

func NewWebhookByIDToken(client *http.Client, logger log.Logger, webhookID string, webhookToken string) (api.Webhook, error) {
	return internal.NewWebhookImpl(client, logger, webhookID, webhookToken), nil
}
