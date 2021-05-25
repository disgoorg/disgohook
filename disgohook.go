package disgohook

import (
	"net/http"
	"strings"

	"github.com/DisgoOrg/disgohook/api"
	"github.com/DisgoOrg/disgohook/internal"
	"github.com/DisgoOrg/log"
)

func NewWebhookByToken(client *http.Client, logger log.Logger, webhookToken string) (api.WebhookClient, error) {
	webhookTokenSplit := strings.SplitN(webhookToken, "/", 2)
	if len(webhookTokenSplit) != 2 {
		return nil, api.ErrMalformedWebhookToken
	}
	return NewWebhookByIDToken(client, logger, api.Snowflake(webhookTokenSplit[0]), webhookTokenSplit[1])
}

func NewWebhookByIDToken(client *http.Client, logger log.Logger, webhookID api.Snowflake, webhookToken string) (api.WebhookClient, error) {
	return internal.NewWebhookClientImpl(client, logger, webhookID, webhookToken), nil
}
