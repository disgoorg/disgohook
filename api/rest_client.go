package api

import (
	"errors"

	"github.com/DisgoOrg/restclient"
)

// Errors when connecting to discord
var (
	ErrBadGateway   = errors.New("bad gateway could not reach discord")
	ErrUnauthorized = errors.New("not authorized for this endpoint")
	ErrBadRequest   = errors.New("bad request please check your request")
	ErrRatelimited  = errors.New("too many requests")
)

// ErrorResponse contains custom errors from discord
type ErrorResponse struct {
	Code    int
	Message string
}

// RestClient is a manager for all of disgohook's HTTP requests
type RestClient interface {
	restclient.RestClient
	WebhookClient() WebhookClient

	GetWebhook(webhookID Snowflake, webhookToken string) (*Webhook, error)
	UpdateWebhook(webhookID Snowflake, webhookToken string, webhookUpdate WebhookUpdate) (*Webhook, error)
	DeleteWebhook(webhookID Snowflake, webhookToken string) error

	CreateWebhookMessage(webhookID Snowflake, webhookToken string, messageCreate WebhookMessageCreate, wait bool, threadID Snowflake) (*WebhookMessage, error)
	UpdateWebhookMessage(webhookID Snowflake, webhookToken string, messageID Snowflake, messageUpdate WebhookMessageUpdate) (*WebhookMessage, error)
	DeleteWebhookMessage(webhookID Snowflake, webhookToken string, messageID Snowflake) error
}
