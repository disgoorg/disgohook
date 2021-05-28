package internal

import (
	"net/http"

	"github.com/DisgoOrg/log"
	"github.com/DisgoOrg/restclient"

	"github.com/DisgoOrg/disgohook/api"
)

var _ api.WebhookClient = (*WebhookClientImpl)(nil)

func NewWebhookClientImpl(client *http.Client, logger log.Logger, id api.Snowflake, token string) api.WebhookClient {
	webhook := &WebhookClientImpl{
		logger:                 logger,
		defaultAllowedMentions: &api.DefaultAllowedMentions,
		id:                     id,
		token:                  token,
	}
	webhook.restClient = newRestClientImpl(client, webhook)
	return webhook
}

type WebhookClientImpl struct {
	restClient             api.RestClient
	logger                 log.Logger
	defaultAllowedMentions *api.AllowedMentions
	id                     api.Snowflake
	token                  string
}

func (h *WebhookClientImpl) RestClient() api.RestClient {
	return h.restClient
}
func (h *WebhookClientImpl) Logger() log.Logger {
	return h.logger
}
func (h *WebhookClientImpl) DefaultAllowedMentions() *api.AllowedMentions {
	return h.defaultAllowedMentions
}
func (h *WebhookClientImpl) SetDefaultAllowedMentions(allowedMentions *api.AllowedMentions) {
	h.defaultAllowedMentions = allowedMentions
}

func (h *WebhookClientImpl) GetWebhook() (*api.Webhook, error) {
	return h.RestClient().GetWebhook(h.id, h.token)
}

func (h *WebhookClientImpl) EditWebhook(webhookUpdate *api.WebhookUpdate) (*api.Webhook, error) {
	return h.RestClient().UpdateWebhook(h.id, h.token, webhookUpdate)
}

func (h *WebhookClientImpl) DeleteWebhook() error {
	return h.RestClient().DeleteWebhook(h.id, h.token)
}

func (h *WebhookClientImpl) SendMessage(message *api.WebhookMessageCreate) (*api.WebhookMessage, error) {
	return h.RestClient().CreateWebhookMessage(h.id, h.token, message, true, "")
}

func (h *WebhookClientImpl) SendContent(content string) (*api.WebhookMessage, error) {
	return h.SendMessage(api.NewWebhookMessageBuilderWithContent(content).Build())
}

func (h *WebhookClientImpl) SendEmbed(embed *api.Embed, embeds ...*api.Embed) (*api.WebhookMessage, error) {
	return h.SendMessage(api.NewWebhookMessageBuilderWithEmbeds(embed, embeds...).Build())
}

func (h *WebhookClientImpl) EditMessage(messageID api.Snowflake, message *api.WebhookMessageUpdate) (*api.WebhookMessage, error) {
	return h.RestClient().UpdateWebhookMessage(h.id, h.token, messageID, message)
}

func (h *WebhookClientImpl) EditContent(messageID api.Snowflake, content string) (*api.WebhookMessage, error) {
	return h.EditMessage(messageID, &api.WebhookMessageUpdate{Content: &content})
}

func (h *WebhookClientImpl) EditEmbed(messageID api.Snowflake, embed *api.Embed, embeds ...*api.Embed) (*api.WebhookMessage, error) {
	return h.EditMessage(messageID, &api.WebhookMessageUpdate{Embeds: append([]*api.Embed{embed}, embeds...)})
}

func (h *WebhookClientImpl) DeleteMessage(messageID api.Snowflake) error {
	return h.RestClient().DeleteWebhookMessage(h.id, h.token, messageID)
}

func (h *WebhookClientImpl) Token() string {
	return h.token
}
func (h *WebhookClientImpl) ID() api.Snowflake {
	return h.id
}
func (h *WebhookClientImpl) URL() string {
	compiledRoute, _ := restclient.GetWebhook.Compile(nil, h.id, h.token)
	return compiledRoute.Route()
}
