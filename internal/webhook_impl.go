package internal

import (
	"net/http"

	"github.com/DisgoOrg/disgohook/api/endpoints"
	"github.com/DisgoOrg/log"

	"github.com/DisgoOrg/disgohook/api"
)

var _ api.Webhook = (*WebhookImpl)(nil)

func NewWebhookImpl(client *http.Client, logger log.Logger, id string, token string) api.Webhook {
	webhook := &WebhookImpl{
		logger:                 logger,
		defaultAllowedMentions: &api.DefaultAllowedMentions,
		id:                     id,
		token:                  token,
	}
	webhook.restClient = newRestClientImpl(client, webhook)
	return webhook
}

type WebhookImpl struct {
	restClient             api.RestClient
	logger                 log.Logger
	defaultAllowedMentions *api.AllowedMentions
	id                     string
	token                  string
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

func (h *WebhookImpl) SendMessage(message *api.WebhookMessageCreate) (*api.WebhookMessage, error) {
	return h.RestClient().CreateWebhookMessage(h.id, h.token, message)
}

func (h *WebhookImpl) SendContent(content string) (*api.WebhookMessage, error) {
	return h.RestClient().CreateWebhookMessage(h.id, h.token, api.NewWebhookMessageBuilderWithContent(content).Build())
}

func (h *WebhookImpl) SendEmbed(embed *api.Embed, embeds ...*api.Embed) (*api.WebhookMessage, error) {
	return h.RestClient().CreateWebhookMessage(h.id, h.token, api.NewWebhookMessageBuilderWithEmbeds(embed, embeds...).Build())
}

func (h *WebhookImpl) EditMessage(messageID string, message *api.WebhookMessageUpdate) (*api.WebhookMessage, error) {
	return h.RestClient().UpdateWebhookMessage(h.id, h.token, messageID, message)
}

func (h *WebhookImpl) EditContent(messageID string, content string) (*api.WebhookMessage, error) {
	return h.RestClient().UpdateWebhookMessage(h.id, h.token, messageID, &api.WebhookMessageUpdate{Content: &content})
}

func (h *WebhookImpl) EditEmbed(messageID string, embed *api.Embed, embeds ...*api.Embed) (*api.WebhookMessage, error) {
	return h.RestClient().UpdateWebhookMessage(h.id, h.token, messageID, &api.WebhookMessageUpdate{Embeds: append([]*api.Embed{embed}, embeds...)})
}

func (h *WebhookImpl) DeleteMessage(messageID string) error {
	return h.RestClient().DeleteWebhookMessage(h.id, h.token, messageID)
}

func (h *WebhookImpl) Token() string {
	return h.token
}
func (h *WebhookImpl) ID() string {
	return h.id
}
func (h *WebhookImpl) URL() string {
	compiledRoute, _ := endpoints.GetWebhook.Compile(h.id, h.token)
	return compiledRoute.Route()
}
