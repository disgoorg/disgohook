package api

import (
	"errors"
	"regexp"

	"github.com/DisgoOrg/log"
)

var WebhookPattern = regexp.MustCompile("(?:https?://)?(?:\\w+\\.)?discord(?:app)?\\.com/api(?:/v\\d+)?/webhooks/(\\d+)/([\\w-]+)(?:/(?:\\w+)?)?")

var ErrMalformedWebhookToken = errors.New("malformed webhook token <id>/<token>")

type Webhook interface {
	RestClient() RestClient
	Logger() log.Logger
	DefaultAllowedMentions() *AllowedMentions
	SetDefaultAllowedMentions(allowedMentions *AllowedMentions)

	SendMessage(message WebhookMessageCreate) (*WebhookMessage, error)
	SendContent(content string) (*WebhookMessage, error)
	SendEmbed(embed *Embed, embeds ...*Embed) (*WebhookMessage, error)

	EditMessage(messageID string, message WebhookMessageUpdate) (*WebhookMessage, error)
	EditContent(messageID string, content string) (*WebhookMessage, error)
	EditEmbed(messageID string, embed *Embed, embeds ...*Embed) (*WebhookMessage, error)

	DeleteMessage(id string) error

	Token() string
	ID() string
	URL() string
}
