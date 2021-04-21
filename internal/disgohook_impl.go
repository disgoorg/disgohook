package internal

import (
	"net/http"

	"github.com/DisgoOrg/disgohook/api"
)

var _ api.DisgoHook = (*DisgoHookImpl)(nil)

func NewDisgoHookImpl(client http.Client, token string, id string) api.DisgoHook {
	return &DisgoHookImpl{
		token:  token,
		id:     id,
		client: client,
	}
}

type DisgoHookImpl struct {
	token  string
	id     string
	client http.Client
}

func (h *DisgoHookImpl) SendMessage() (*api.Message, error) {
	return nil, nil
}
func (h *DisgoHookImpl) Token() string {
	return ""
}
func (h *DisgoHookImpl) ID() string {
	return ""
}
