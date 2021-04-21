package api

import (
	"errors"
	"regexp"
)

var WebhookPattern = regexp.MustCompile("(?:https?://)?(?:\\w+\\.)?discord(?:app)?\\.com/api(?:/v\\d+)?/webhooks/(\\d+)/([\\w-]+)(?:/(?:\\w+)?)?")

var ErrMalformedWebhookToken = errors.New("malformed webhook token <id>/<token>")

type DisgoHook interface {
	RestClient() RestClient
	SendMessage() (*Message, error)
	Token() string
	ID() string
}
