package disgohook

import (
	"strings"

	"github.com/DisgoOrg/disgohook/api"
	"github.com/DisgoOrg/disgohook/internal"
)

func NewDisgoHookByToken(webhookToken string) (api.DisgoHook, error) {
	webhookTokenSplit := strings.SplitN(webhookToken, "/", 2)
	if len(webhookTokenSplit) != 2 {
		return nil, api.ErrMalformedWebhookToken
	}
	return internal.NewDisgoHookImpl(webhookTokenSplit[0], webhookTokenSplit[1]), nil
}

func NewDisgoHookByID(id string, token string) (api.DisgoHook, error) {
	return nil, nil
}
