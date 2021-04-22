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
	SendContent(message string) (*WebhookMessage, error)
	SendEmbed(embed *Embed, embeds ...*Embed)

	EditMessage(id string, message WebhookMessageUpdate) (*WebhookMessage, error)
	EditContent(message string) (*WebhookMessage, error)
	EditEmbed(embed *Embed, embeds ...*Embed)

	DeleteMessage(id string) error

	Token() string
	ID() string
	URL() string
}
