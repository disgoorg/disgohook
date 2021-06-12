package api

import (
	"errors"
	"regexp"

	"github.com/DisgoOrg/log"
)

// WebhookURLPattern is the URL pattern for Webhook(s)
var WebhookURLPattern = regexp.MustCompile("(?:https?://)?(?:\\w+\\.)?discord(?:app)?\\.com/api(?:/v\\d+)?/webhooks/(\\d+)/([\\w-]+)(?:/(?:\\w+)?)?")

// ErrMalformedWebhookToken is returned for invalid webhook tokens
var ErrMalformedWebhookToken = errors.New("malformed webhook token <id>/<token>")

// WebhookClient lets you edit/send WebhookMessage(s) or update/delete the Webhook
type WebhookClient interface {
	RestClient() RestClient
	Logger() log.Logger
	DefaultAllowedMentions() *AllowedMentions
	SetDefaultAllowedMentions(allowedMentions *AllowedMentions)

	GetWebhook() (*Webhook, error)
	EditWebhook(webhookUpdate WebhookUpdate) (*Webhook, error)
	DeleteWebhook() error

	SendMessage(message WebhookMessageCreate) (*WebhookMessage, error)
	SendContent(content string) (*WebhookMessage, error)
	SendEmbed(embeds ...Embed) (*WebhookMessage, error)

	EditMessage(messageID Snowflake, message WebhookMessageUpdate) (*WebhookMessage, error)
	EditContent(messageID Snowflake, content string) (*WebhookMessage, error)
	EditEmbed(messageID Snowflake, embeds ...Embed) (*WebhookMessage, error)

	DeleteMessage(id Snowflake) error

	Token() string
	ID() Snowflake
	URL() string
}
