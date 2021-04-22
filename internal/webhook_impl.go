package internal

import (
	"github.com/DisgoOrg/log"
	"net/http"

	"github.com/DisgoOrg/disgohook/api"
)

//var _ api.Webhook = (*WebhookImpl)(nil)

func NewWebhookImpl(client *http.Client, logger log.Logger, token string, id string) *WebhookImpl {
	webhook := &WebhookImpl{
		logger:                 logger,
		defaultAllowedMentions: &api.DefaultAllowedMentions,
		token:                  token,
		id:                     id,
	}
	webhook.restClient = newRestClientImpl(client, webhook)
	return webhook
}

type WebhookImpl struct {
	restClient             api.RestClient
	logger                 log.Logger
	defaultAllowedMentions *api.AllowedMentions
	token                  string
	id                     string
}

func (h *WebhookImpl) RestClient() api.RestClient {
	return h.restClient
}
func (h *WebhookImpl) Logger() log.Logger {
	return h.logger
}
func (h *WebhookImpl) DefaultAllowedMentions() *api.AllowedMentions {
	return h.defaultAllowedMentions
}
func (h *WebhookImpl) SetDefaultAllowedMentions(allowedMentions *api.AllowedMentions) {
	h.defaultAllowedMentions = allowedMentions
}


func (h *WebhookImpl) Token() string {
	return h.token
}
func (h *WebhookImpl) ID() string {
	return h.id
}
