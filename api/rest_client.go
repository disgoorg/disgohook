package api

import (
	"errors"

	"github.com/DisgoOrg/disgohook/api/endpoints"
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

type RestClient interface {
	UserAgent() string
	Request(route endpoints.CompiledAPIRoute, rqBody interface{}, rsBody interface{}) error

	CreateWebhookMessage(webhookID string, webhookToken string, message *WebhookMessageCreate) (*WebhookMessage, error)
	UpdateWebhookMessage(webhookID string, webhookToken string, messageID string, message *WebhookMessageUpdate) (*WebhookMessage, error)
	DeleteWebhookMessage(webhookID string, webhookToken string, messageID string) error
}
