package api

import "github.com/DisgoOrg/disgohook/api/endpoints"

type RestClient interface {
	UserAgent() string
	Request(route endpoints.CompiledAPIRoute, rqBody interface{}, rsBody interface{}) error

	GetWebhook(webhookID string, webhookToken string) (*Webhook, error)
	UpdateWebhook(webhookID string, webhookToken string, webhookUpdate UpdateWebhook) (*Webhook, error)
	DeleteWebhook(webhookID string, webhookToken string) error

	CreateWebhookMessage(webhookID string, webhookToken string) (*Message, error)
	CreateWebhookMessageSlack(webhookID string, webhookToken string) (*Message, error)
	CreateWebhookMessageGithub(webhookID string, webhookToken string) (*Message, error)
	UpdateWebhookMessage(webhookID string, webhookToken string) (*Message, error)
	DeleteWebhookMessage(webhookID string, webhookToken string) error
}
